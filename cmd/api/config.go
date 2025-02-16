package main

import (
	"github.com/KengoWada/go-todos/internal/store"
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
