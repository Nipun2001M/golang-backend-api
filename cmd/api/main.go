package main

import (
	"fmt"
	"log"
)
import "github.com/Nipun2001M/golang-backend-api/internal/env"
import "github.com/Nipun2001M/golang-backend-api/internal/store"
import "github.com/Nipun2001M/golang-backend-api/internal/db"

func main() {
	cfg := config{
		addresss: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5433/gosocial?sslmode=disable"),
			maxOpenConn: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConn: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdletime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	fmt.Println("Attempting to connect with:", cfg.db.addr)
	dataB, err := db.New(cfg.db.addr, cfg.db.maxOpenConn, cfg.db.maxIdleConn, cfg.db.maxIdletime)

	if err != nil {
		log.Panic(err)
		fmt.Print(err)
	}

	defer dataB.Close()
	storedb := store.NewStorage(dataB)

	log.Printf("DB Connection established")

	app := appication{
		config: cfg,
		store:  storedb,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
