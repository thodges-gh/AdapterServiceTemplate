package main

// Chainlink contains the fields necessary for receiving the data
// that will be sent to the adapter and for returing back to the
// node itself.
type Chainlink struct {
	ID      string `json:"id"`
	Data    Data   `json:"data"`
	Error   string `json:"error,omitempty"`
	Pending bool   `json:"pending,omitempty"`
}

// Data should be modeled after the target endpoint. You can use
// something like JSON-to-Go to make this easy:
// https://mholt.github.io/json-to-go/
type Data struct {
	Value string `json:"value"`
	Last  string `json:"last"`
	Other string `json:"other"`
}

// GetData is where you woud reach out to your desired endpoint and
// store the response in the Data struct.
func GetData(cl Chainlink) Chainlink {
	cl.Data = Data{
		Value: "true",
		Last:  "1111",
		Other: "crypto",
	}

	// The Pending field is optional
	//cl.Pending = false

	return cl
}
