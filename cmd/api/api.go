package main

import (
	"github.com/Nipun2001M/golang-backend-api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type appication struct {
	config config
	store  store.Storage
}

type config struct {
	addresss string
	db       dbConfig
}

type dbConfig struct {
	addr        string
	maxOpenConn int
	maxIdleConn int
	maxIdletime string
}

func (app *appication) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r

}

func (app *appication) run(mux *chi.Mux) error {
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
