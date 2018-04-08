package main

import (
	null "gopkg.in/guregu/null.v3"
)
// Chainlink contains the fields necessary for receiving the data
// that will be sent to the adapter and for returing back to the
// node itself.
type Chainlink struct {
	ID      string `json:"id"`
	Data    Data   `json:"data"`
}

type RunResult struct {
	JobRunID     string      `json:"jobRunId"`
	Data         Data        `json:"data"`
	Status       string      `json:"status"`
	ErrorMessage null.String `json:"error"`
	Pending      bool        `json:"pending"`
}

// Data should be modeled after the target endpoint. You can use
// something like JSON-to-Go to make this easy:
// https://mholt.github.io/json-to-go/
type Data struct {
	Value string `json:"value,omitempty"`
	Last  string `json:"last,omitempty"`
	Other string `json:"other,omitempty"`
}

// GetData is where you woud reach out to your desired endpoint and
// store the response in the Data struct.
func GetData(cl Chainlink) Chainlink {
	rr := RunResult{
		JobRunID: cl.ID,
		Data : Data{
			Value: "SomeValue",
			Last:  "1111",
			Other: "GetData",
		},
		Status: "completed",
		Pending: false,
	}

	return rr
}

// GetPending returns the pending field as true without data.
func GetPending(cl Chainlink) Chainlink {
	rr := RunResult{
		JobRunID: cl.ID,
		Pending: true,
	}

	return rr
}

// GetBigInt returns the max value of a uInt256 type in Solidity.
func GetBigInt(cl Chainlink) RunResult {
	rr := RunResult{
		JobRunID: cl.ID,
		Data : Data{
			Value: "115792089237316195423570985008687907853269984665640564039457584007913129639935",
			Last:  "1111",
			Other: "GetBigInt",
		},
		Status: "completed",
		Pending: false,
	}

	return rr
}
