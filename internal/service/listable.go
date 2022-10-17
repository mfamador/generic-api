// Package service contains the business logic for services
package service

import (
	"genericsapi/internal/genericsapiv1"
	"genericsapi/internal/repository"
)

// Listable exposes the interface to do operations on the generic entity
type Listable[T any] interface {
	List(uint, *string, []*genericsapiv1.Filter) ([]*T, *string, error)
}

type listable[T any] struct {
	repository repository.ListRepository[T]
}

// NewListable creates a generic Listable service
func NewListable[T any](repo repository.ListRepository[T]) Listable[T] {
	return &listable[T]{
		repository: repo,
	}
}

// List retrieves a list of listable entities
func (s *listable[T]) List(limit uint, cursor *string, filter []*genericsapiv1.Filter) (res []*T, c *string, err error) {
	sn, span, traceID := startSpan(s)
	logRequest(traceID, sn, cursor, filter)
	if verr := validateFilter(filter); verr != nil {
		return nil, nil, verr
	}
	res, c, err = s.repository.List(limit, cursor, filter)
	logResponse(traceID, sn, cursor, err)
	span.Finish()
	return res, c, err
}
