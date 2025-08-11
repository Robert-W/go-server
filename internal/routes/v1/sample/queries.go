package sample

import "time"

type Sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func listSamplesQuery() (*[]Sample, error) {
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

	return &samples, nil
}

func createSamplesQuery() (*[]Sample, error) {
	samples := []Sample{
		{
			Id:        "111",
			Value:     "New Sample",
			Timestamp: time.Now(),
		},
	}

	return &samples, nil
}

func readSampleQuery() (*Sample, error) {
	sample := Sample{
		Id:        "123",
		Value:     "Sample Read",
		Timestamp: time.Now(),
	}

	return &sample, nil
}

func updateSampleQuery() (*Sample, error) {
	sample := Sample{
		Id:        "321",
		Value:     "Sample Update",
		Timestamp: time.Now(),
	}

	return &sample, nil
}

func deleteSampleQuery() error {
	return nil
}
