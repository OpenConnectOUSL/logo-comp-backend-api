package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router{
	// initialize a new httprouter instance
	router := httprouter.New()

	// Register the relevant URL methods
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.createUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.showUserHandler)
	
	return router
}