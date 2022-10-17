// Package service contains the business logic for ReadingsLastValues
package service //nolint

import (
	"genericsapi/internal/repository"
)

// Status exposes the interface to do operations on the Status entity
type Status interface {
	CheckStatus() (*string, error)
}

type status struct {
	repository repository.Status
}

// NewStatus creates a Status service
func NewStatus(repo repository.Status) Status {
	return &status{
		repo,
	}
}

// CheckStatus check the status
func (s *status) CheckStatus() (*string, error) {
	return s.repository.CheckStatus()
}
