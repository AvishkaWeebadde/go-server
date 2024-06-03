package main

import (
  "log"
  "net/http"

  "github.com/AvishkaWeebadde/go-server/api"
  "github.com/AvishkaWeebadde/go-server/middleware"
)

func main() {
  const filepathRoot = "."
  const port = "8080"

  // Move apiConfig to middleware package
  apiCfg := middleware.GetApiConfig()

  mux := api.NewRouter()
  mux.Handle("/app/*", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
  mux.HandleRoutes()

  srv := &http.Server{
    Addr:    ":" + port,
    Handler: mux,
  }

  log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
  log.Fatal(srv.ListenAndServe())
}
