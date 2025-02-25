package main

import (
	"log"

	"github.com/amrremam/Golang-RestAPIs.git/internal/store"
)


func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr: "DB_addr","postgres://",
		},
	}
	
	store := store.NewStorage(nil)
	
	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}