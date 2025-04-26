package main

import (
	"log"
	"net/http"
)

type appication struct {
	config config
}

type config struct {
	addresss string
}

func (app *appication) run() error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    app.config.addresss,
		Handler: mux,
	}
	log.Printf("Server has started at %s", app.config.addresss)
	return srv.ListenAndServe()

}
