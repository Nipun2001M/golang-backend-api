package main

import (
	"log"
	"net/http"
	"time"
)

type appication struct {
	config config
}

type config struct {
	addresss string
}

func (app *appication) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	return mux

}

func (app *appication) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addresss,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at %s", app.config.addresss)
	return srv.ListenAndServe()

}
