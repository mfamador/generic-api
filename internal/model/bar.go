// Package model contains the Packet model
package model

import "time"

// Bar defines the bar entity
type Bar struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Value       float64   `json:"value,omitempty"`
	SpecificBar string    `json:"specific_bar,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}
