package main

import (
	"log"

	"github.com/KengoWada/go-todos/internal/database"
	"github.com/KengoWada/go-todos/internal/env"
	"github.com/KengoWada/go-todos/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr:        env.GetString("ADDR", ":8000"),
		environment: env.GetString("ENVIRONMENT", "production"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost/todos?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := database.New(
		cfg.dbConfig.addr,
		cfg.dbConfig.maxOpenConns,
		cfg.dbConfig.maxIdleConns,
		cfg.dbConfig.maxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(db)

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
