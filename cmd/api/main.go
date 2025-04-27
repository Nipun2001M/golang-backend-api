package main

import (
	"log"
)
import "github.com/Nipun2001M/golang-backend-api/internal/env"

func main() {
	cfg := config{
		addresss: env.GetString("ADDR", ":8080"),
	}

	app := appication{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
