package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "TaskRun",
		Method:      "POST",
		Pattern:     "/",
		HandlerFunc: TaskRun,
	},
	{
		Name:        "PendingTaskRun",
		Method:      "POST",
		Pattern:     "/pending",
		HandlerFunc: PendingTaskRun,
	},
	{
		Name:        "ReturnBigInt",
		Method:      "POST",
		Pattern:     "/big",
		HandlerFunc: ReturnBigInt,
	},
	{
		Name:        "InputDataExample",
		Method:      "POST",
		Pattern:     "/input",
		HandlerFunc: InputDataExample,
	},
	{
		Name:        "Resume",
		Method:      "POST",
		Pattern:     "/resume",
		HandlerFunc: ResumeFromPending,
	},
	{
		Name:        "Error",
		Method:      "POST",
		Pattern:     "/error",
		HandlerFunc: ReturnError,
	},
	// Add more routes here if needed, keep in
	// mind that Chainlink will POST to them for
	// task runs.
}

// https://github.com/gorilla/mux#middleware
type MiddlewareFunc func(http.Handler) http.Handler

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
