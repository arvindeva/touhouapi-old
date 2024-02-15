package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter" // New import
	"github.com/justinas/alice"
)

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w, r)
	})

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the Touhou API Project")
	})
	router.HandlerFunc(http.MethodGet, "/touhou", app.getTouhous)
	router.HandlerFunc(http.MethodGet, "/touhou/:id", app.getTouhouByID)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
