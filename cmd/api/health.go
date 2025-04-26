package main

import "net/http"

func (app *appication) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
