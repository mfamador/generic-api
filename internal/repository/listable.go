// Package repository describes the interface of a Packet datasource
package repository

import (
	"genericsapi/internal/genericsapiv1"
)

// Listable exposes the interface to access a data source
type Listable[T any] interface {
	List(uint, *string, []*genericsapiv1.Filter) ([]*T, *string, error)
}
