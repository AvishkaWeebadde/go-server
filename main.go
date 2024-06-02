package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//add prefix /app/ to all file serving routes
	router.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("./"))))

	router.Handle("/", http.FileServer(http.Dir("./")))

	fs := http.FileServer(http.Dir("./assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router.HandleFunc("/healthz", handlerHealthz)

	log.Printf("Server started at localhost%s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
