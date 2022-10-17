// Package genericsapi has the implementation of generics API services
package genericsapi

import (
	"context"
	"fmt"
	"genericsapi/internal/datasource/cassandra"
	"genericsapi/internal/genericsapiv1"
	"genericsapi/internal/model"
	"genericsapi/internal/service"

	"github.com/gocql/gocql"
)

// Foo exposes the interface
type Foo interface {
	Read(context.Context, *genericsapiv1.ReadRequest) (*genericsapiv1.ReadFooReply, error)
}

type foo struct {
	service service.Listable[model.Foo]
}

// NewFoo instantiates a new service
func NewFoo(session *gocql.Session) Foo {
	return &foo{
		service.NewListable(cassandra.NewFoo(session)),
	}
}

// GetReadings retrieves the readings from cassandra
func (d *foo) Read(ctx context.Context, request *genericsapiv1.ReadRequest) (*genericsapiv1.ReadFooReply, error) {
	res, cs, err := d.service.List(uint(request.Limit), func() *string {
		if len(request.Cursor) > 0 {
			return &request.Cursor
		}
		return nil
	}(), request.Filters)
	if err != nil {
		return nil, fmt.Errorf("error retrieving values: %v", err)
	}
	return mapFoo(cs, res)
}
