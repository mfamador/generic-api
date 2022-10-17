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

// Bar exposes the interface
type Bar interface {
	Read(context.Context, *genericsapiv1.ReadRequest) (*genericsapiv1.ReadBarReply, error)
}

type bar struct {
	service service.Listable[model.Bar]
}

// NewBar instantiates a new service
func NewBar(session *gocql.Session) Bar {
	return &bar{
		service.NewListable(cassandra.NewBar(session)),
	}
}

// GetReadings retrieves the readings from cassandra
func (d *bar) Read(ctx context.Context, request *genericsapiv1.ReadRequest) (*genericsapiv1.ReadBarReply, error) {
	res, cs, err := d.service.List(uint(request.Limit), func() *string {
		if len(request.Cursor) > 0 {
			return &request.Cursor
		}
		return nil
	}(), request.Filters)
	if err != nil {
		return nil, fmt.Errorf("error retrieving values: %v", err)
	}
	return mapBar(cs, res)
}
