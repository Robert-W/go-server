package sample

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func ListSamples(res http.ResponseWriter, _ *http.Request) {
	samples := []Sample{
		{
			Id:        "111",
			Value:     "First Sample",
			Timestamp: time.Now(),
		},
		{
			Id:        "222",
			Value:     "Second Sample",
			Timestamp: time.Now(),
		},
		{
			Id:        "333",
			Value:     "Third Sample",
			Timestamp: time.Now(),
		},
	}

	sampleJson, err := json.Marshal(samples)

	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}
