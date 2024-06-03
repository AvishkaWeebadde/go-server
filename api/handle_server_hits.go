package api

import (
	"fmt"
	"net/http"

	"github.com/AvishkaWeebadde/go-server/middleware"
)

func handleFilseServerHitsReset(w http.ResponseWriter, r *http.Request) {
	middleware.GetApiConfig().ResetFileServerHits()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", middleware.GetApiConfig().GetFileServerHits())))
}
