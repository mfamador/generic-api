// Package genericsapi has the implementation of generics API services
package genericsapi

import (
	"context"
	"fmt"
	"genericsapi/internal/genericsapiv1"
	"net/http"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// Status exposes the interface to the health service
type Status interface {
	Ready(echo.Context) error
	Ping(echo.Context) error
}

type status struct {
	ctx    context.Context
	client genericsapiv1.StatusServiceClient
}

// NewStatus creates the status struct
func NewStatus(port int) (Status, error) {
	ctx := context.Background()

	// Set up a connection to the server.
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	d := genericsapiv1.NewStatusServiceClient(conn)
	return &status{
		ctx:    ctx,
		client: d,
	}, nil
}

func (hs *status) Ready(c echo.Context) error {
	_, err := hs.client.Status(hs.ctx, &genericsapiv1.StatusRequest{})
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusServiceUnavailable)
	}

	return c.String(http.StatusOK, "OK")
}

func (hs *status) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
