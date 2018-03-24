package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// TaskRun reads the input JSON given from the core, calls the
// GetData method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func TaskRun(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Log the input
	log.Println("Input:")
	log.Println(string(body))

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := GetData(cl)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	// Log the output
	outString, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	log.Println("Output:")
	log.Println(string(outString))
}
