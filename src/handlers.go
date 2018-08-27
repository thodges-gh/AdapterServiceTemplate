package main

import (
	"bytes"
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

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

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
	w.Header().Set("Authorization", "b3a2064827904ffb9c45567e3b04d382")
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

// PendingTaskRun reads the input JSON given from the core, calls the
// GetPending method, and returns a RunResult with the pending status
// set to true. It will log both the input and output JSON for troubleshooting.
func PendingTaskRun(w http.ResponseWriter, r *http.Request) {
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

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := GetPending(cl)

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

// ResumeFromPending allows you to resume a pending run, given its
// JobRunID.
func ResumeFromPending(w http.ResponseWriter, r *http.Request) {
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

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := GetData(cl)
	url := "http://localhost:6688/v2/runs/" + cl.ID
	outString, err := json.Marshal(result)
	// trimmed := outString[1:len(outString)-1]
	if err != nil {
		panic(err)
	}

	log.Println("PATCHing to:")
	log.Println(url)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(outString))
	req.Header.Set("Authorization", "Bearer 7a76ab15fd984d4ba6171c064f537c29")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(responseBody))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(req.Body); err != nil {
		panic(err)
	}

	// Log the output
	
	log.Println("Output:")
	log.Println(string(outString))
}

// ReturnBigInt reads the input JSON given from the core, calls the
// GetBigInt method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func ReturnBigInt(w http.ResponseWriter, r *http.Request) {
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

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := GetBigInt(cl)

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

// InputDataExample reads the input JSON given from the core, calls the
// GetRestData method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func InputDataExample(w http.ResponseWriter, r *http.Request) {
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

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := GetInputData(cl)

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
