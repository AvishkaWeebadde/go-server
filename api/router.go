package api

import "net/http"

type router struct {
	*http.ServeMux
}

func NewRouter() *router {
	return &router{http.NewServeMux()}
}

func prefixHandler(prefix string, handler http.Handler) http.Handler {
	return http.StripPrefix(prefix, handler)
}

func (r *router) HandleRoutes() {
	r.HandleFunc("/healthz", handlerHealthz)
	r.HandleFunc("/reset", handleFilseServerHitsReset)
	r.HandleFunc("/metrics", handleMetrics)
}