package main

import "log"

func main() {
	cfg := config{
		addresss: ":8080",
	}

	app := appication{
		config: cfg,
	}

	log.Fatal(app.run())
}
