// Package server contains the common HTTP server logic, e.g.:
//
//	routes, validators, error handling
package server

import (
	"fmt"
	"genericsapi/internal/datasource/cassandra"
	"genericsapi/internal/genericsapi"
	"genericsapi/internal/genericsapiv1"
	"github.com/labstack/echo/v4"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Config defines the handler configuration
type Config struct {
	GrpcPort            int    `yaml:"grpcPort"`
	RestPort            int    `yaml:"restPort"`
	TracingAgentAddress string `yaml:"tracingAgentAddress"`
}

// Start runs the gRPC server
func Start(grpcServer *grpc.Server, list net.Listener) error {
	log.Info().Msg("Data API gRPC server")
	return grpcServer.Serve(list)
}

// GetGRPCServer returns the gRPC server
func GetGRPCServer(conf Config, confC *cassandra.Config) (*grpc.Server, net.Listener, error) {
	const (
		fooKeyspace = "foo"
		barKeyspace = "bar"
	)
	fooSession, err := cassandra.GetSession(fooKeyspace, confC)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get foo session: %v", err)
	}
	barSession, err := cassandra.GetSession(barKeyspace, confC)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get bar session: %v", err)
	}
	grpcServer := grpc.NewServer()
	genericsapiv1.RegisterFooServiceServer(grpcServer, genericsapi.NewFoo(fooSession))
	genericsapiv1.RegisterBarServiceServer(grpcServer, genericsapi.NewBar(barSession))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		return nil, nil, fmt.Errorf("error creating listener: %v", err)
	}
	return grpcServer, listener, nil
}

// RunApp runs the API
func RunApp(sConf Config, confC *cassandra.Config) error {
	grpcServer, lis, err := GetGRPCServer(sConf, confC)
	if err != nil {
		return fmt.Errorf("failed to build grpcServer: %v", err)
	}
	go func() {
		if e := Start(grpcServer, lis); e != nil {
			log.Fatal().Msgf("Failed to start the gRPC server")
		}
	}()
	e := echo.New()
	hs, err := genericsapi.NewStatus(sConf.GrpcPort)
	if err != nil {
		return fmt.Errorf("failed to start the gRPC client for health checks server: %v", err)
	}
	e.GET("/ready", hs.Ready)
	e.GET("/ping", hs.Ping)
	if err := e.Start(fmt.Sprintf(":%d", sConf.RestPort)); err != nil {
		return fmt.Errorf("failed to start the HTTP server: %v", err)
	}
	return nil
}
