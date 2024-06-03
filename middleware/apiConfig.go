package middleware

import (
  "log"
  "net/http"
  "strings"
  "sync"
)

type apiConfig struct {
  fileServerHits int
  mu             sync.Mutex
}

var (
  instance *apiConfig
  once     sync.Once
)

func GetApiConfig() *apiConfig {
  once.Do(func() {
    instance = &apiConfig{}
  })
  return instance
}

func (a *apiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !shouldExcludePath(r.URL.Path) {
      a.mu.Lock()
      a.fileServerHits++
      a.mu.Unlock()
      log.Printf("Incremented hit count for path: %s", r.URL.Path)
    } else {
      log.Printf("Excluded path from hit count: %s", r.URL.Path)
    }
    w.Header().Set("Cache-Control", "no-cache")
    next.ServeHTTP(w, r)
  })
}

func shouldExcludePath(path string) bool {
  excludePaths := []string{"/metrics", "/healthz", "/reset", "/favicon.ico", "/robots.txt"}
  staticExtensions := []string{".css", ".js", ".jpg", ".jpeg", ".gif", ".svg", ".ico"}

  for _, p := range excludePaths {
    if strings.HasPrefix(path, p) {
      return true
    }
  }

  for _, ext := range staticExtensions {
    if strings.HasSuffix(path, ext) {
      return true
    }
  }

  // Check if path starts with "/app/" and has a static file extension
  if strings.HasPrefix(path, "/app/") {
    for _, ext := range staticExtensions {
      if strings.HasSuffix(path, ext) {
        return true
      }
    }
  }

  return false
}

func (a *apiConfig) ResetFileServerHits() {
  a.mu.Lock()
  defer a.mu.Unlock()
  a.fileServerHits = 0
}

func (a *apiConfig) GetFileServerHits() int {
  a.mu.Lock()
  defer a.mu.Unlock()
  return a.fileServerHits
}