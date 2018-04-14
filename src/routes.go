package main

import (
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
		Name:        "RESTExample",
		Method:      "POST",
		Pattern:     "/rest/{other}",
		HandlerFunc: RestExample,
	},
	// Add more routes here if needed, keep in
	// mind that Chainlink will POST to them for
	// task runs.
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
