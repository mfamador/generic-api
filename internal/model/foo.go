// Package model contains the Packet model
package model

import "time"

// Foo defines the foo entity
type Foo struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Value       float64   `json:"value,omitempty"`
	SpecificFoo string    `json:"specific_foo,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}
