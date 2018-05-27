package main

import (
	null "gopkg.in/guregu/null.v3"
)

// Chainlink contains the fields necessary for receiving the data
// that will be sent to the adapter and for returing back to the
// node itself.
type Chainlink struct {
	ID   string `json:"id"`
	Data Data   `json:"data"`
}

// RunResult formats the response data to the format that the node
// expects.
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
	Value   string  `json:"value,omitempty"`
	Details Details `json:"details,omitempty"`
	Other   string  `json:"other,omitempty"`
}

// Details is just a simple nested JSON object.
type Details struct {
	Close   string `json:"close,omitempty"`
	Open    string `json:"open,omitempty"`
	Current string `json:"current,omitempty"`
}

// GetData is where you woud reach out to your desired endpoint and
// store the response in the Data struct.
func GetData(cl Chainlink) RunResult {
	rr := RunResult{
		JobRunID: cl.ID,
		Data: Data{
			Value: "SomeValue",
			Details: Details{
				Close:   "100",
				Open:    "110",
				Current: "111",
			},
			Other: "GetData",
		},
		Status:  "completed",
		Pending: false,
	}

	return rr
}

// GetPending returns the pending field as true without data.
func GetPending(cl Chainlink) RunResult {
	rr := RunResult{
		JobRunID: cl.ID,
		Pending:  true,
	}

	return rr
}

// GetBigInt returns the max value of a uInt256 type in Solidity.
func GetBigInt(cl Chainlink) RunResult {
	rr := RunResult{
		JobRunID: cl.ID,
		Data: Data{
			Value: "115792089237316195423570985008687907853269984665640564039457584007913129639934",
			Details: Details{
				Close:   "100",
				Open:    "110",
				Current: "111",
			},
			Other: "GetBigInt",
		},
		Status:  "completed",
		Pending: false,
	}

	return rr
}

// GetInputData allows data to be passed into the adapter and the result
// will be determined based on what it is given.
// Pass in '{"other": "GetData"}' to retrieve the value "10000"
// Pass in '{"other": "GetBigInt"}' to retrieve the value "20000"
// Pass in '{"other": "GetRestData"}' to retrieve the value "30000"
func GetInputData(cl Chainlink) RunResult {
	params := cl.Data
	datas := []Data{
		Data{
			Value: "10000",
			Details: Details{
				Close:   "100",
				Open:    "110",
				Current: "111",
			},
			Other: "GetData",
		},
		Data{
			Value: "20000",
			Details: Details{
				Close:   "200",
				Open:    "210",
				Current: "211",
			},
			Other: "GetBigInt",
		},
		Data{
			Value: "30000",
			Details: Details{
				Close:   "300",
				Open:    "310",
				Current: "311",
			},
			Other: "GetRestData",
		}}

	rr := RunResult{
		JobRunID: cl.ID,
		Status:   "completed",
		Pending:  false,
	}

	for _, item := range datas {
		if item.Other == params.Other {
			rr.Data = item
			break
		}
	}

	return rr
}

// GetReportData allows data to be passed into the adapter and the result
// will be determined based on what it is given.
func GetReportData(cl Chainlink) RunResult {
	params := cl.Data
	datas := []Data{
		Data{
			Value: "10000",
			Details: Details{
				Close:   "100",
				Open:    "110",
				Current: "111",
			},
			Other: "GetData",
		},
		Data{
			Value: "20000",
			Details: Details{
				Close:   "200",
				Open:    "210",
				Current: "211",
			},
			Other: "GetBigInt",
		},
		Data{
			Value: "30000",
			Details: Details{
				Close:   "300",
				Open:    "310",
				Current: "311",
			},
			Other: "GetRestData",
		}}

	rr := RunResult{
		JobRunID: cl.ID,
		Status:   "completed",
		Pending:  false,
	}

	for _, item := range datas {
		if item.Other == params.Other {
			rr.Data = item
			break
		}
	}

	return rr
}
