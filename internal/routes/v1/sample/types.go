package sample

import "time"

type sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
