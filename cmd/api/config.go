package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KengoWada/go-todos/internal/store"
	"github.com/KengoWada/go-todos/internal/utils"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr        string
	environment string
	dbConfig    dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			response := struct {
				Message string `json:"message"`
			}{Message: "welcome"}
			utils.WriteJSONResponse(w, http.StatusOK, response)
		})
	})

	return mux
}

func (app *application) run(mux http.Handler) error {
	svr := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started. http://localhost%s", app.config.addr)
	return svr.ListenAndServe()
}
