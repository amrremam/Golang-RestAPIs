package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amrremam/Golang-RestAPIs.git/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type application struct{
	config config
	store store.Storage
	db dbConfig
}


type config struct{
	addr string
}

type dbConfig struct{
	addr string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime	 string
}

func (app *application) mount() http.Handler{
	r := chi.NewRouter()
	
	// use recoverer to avoid panics
	r.Use(middleware.Recoverer)

	r.Use(middleware.Logger)
	
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}
	log.Println("Server is started")
	return srv.ListenAndServe()
}