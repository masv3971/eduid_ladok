package model

import "time"

var (
	// StatusOK ok
	StatusOK = "STATUS_OK"
	// StatusNotOK not ok
	StatusNotOK = "Not OK"
)

// Status type
type Status struct {
	Name       string    `json:"name"`
	SchoolName string    `json:"school_name"`
	Healthy    bool      `json:"healsthy"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}
