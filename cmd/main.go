package main

import (
	"fmt"
	"genericsapi/internal/config"
	"genericsapi/internal/server"
	"genericsapi/internal/tracing"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Golang gRPC generics API")
	jaeger := tracing.NewJaeger(config.Config.Tracing)
	defer jaeger.Close()
	if err := server.RunApp(config.Config.Server, &config.Config.Cassandra); err != nil {
		log.Panic().Msgf("failed to run app: %v", err)
	}
}
