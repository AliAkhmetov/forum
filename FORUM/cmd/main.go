package main

import (
	"log"

	"forum/repository"
	"forum/server"
)

const (
	newDbName       = "./st.db"
	initSqlFileName = "./init-up.sql"
)

func main() {
	// New store instance
	storage, err := repository.New(newDbName)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	// Init DB by init-up.sql
	if err := storage.Init(initSqlFileName); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	// New Кepository struct with interfaces
	repos := repository.NewRepository(storage.Db)

	handler := server.NewHandler(repos)
	server.Server(handler)
}
