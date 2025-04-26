package main

import "log"

func main() {
	cfg := config{
		addresss: ":8080",
	}

	app := appication{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
