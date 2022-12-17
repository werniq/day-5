package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.succededPayment)
	mux.Get("/", app.Home)
	mux.Get("/widget/{id}", app.ChargeOnce)

	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
