package middleware

import (
	"log"
	"net/http"
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

		a.mu.Lock()
		a.fileServerHits++
		a.mu.Unlock()
		log.Printf("Incremented hit count for path: %s", r.URL.Path)
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
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
