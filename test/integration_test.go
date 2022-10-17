// test package for the application
package test

import (
	"context"
	"flag"
	"fmt"
	"genericsapi/internal/genericsapiv1"
	"genericsapi/internal/server"
	"genericsapi/internal/tracing"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/cucumber/godog/colors"

	"github.com/rs/zerolog"

	. "genericsapi/internal/config"
	. "genericsapi/internal/test"

	"github.com/cucumber/godog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

var (
	w    *World
	tt   testing.T
	opts = godog.Options{Output: colors.Colored(os.Stdout)}
)

func TestIntegration(t *testing.T) {
	tt = *t //nolint
}

// Run runs the test suite
func Run(m *testing.M, testSuite godog.TestSuite) {
	flag.Parse()
	opts.Paths = flag.Args()
	os.Args = []string{""}
	testSuite.Options = &opts
	status := testSuite.Run()
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	flag.Parse()
	if testing.Short() {
		log.Warn().Msg("skipping integration tests")
		return
	}
	log.Info().Msg("starting integration tests")
	Run(m, godog.TestSuite{
		Name:                 "CTM Ingestor",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	})
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	jaeger := tracing.NewJaeger(Config.Tracing)
	// start containers
	dockerEnv, err := containers.NewDockerEnvironment()
	if err != nil {
		log.Error().Msgf("error starting docker env: %v", err)
		return
	}
	dockerEnv.SetExpire(240)
	dockerEnv.SetMaxWait(4 * time.Minute)
	useLocalContainers := false
	if l, perr := strconv.ParseBool(os.Getenv("TESTS_LOCAL_CONTAINERS")); perr == nil {
		useLocalContainers = l
	}
	c, err := dockerEnv.StartContainers(func() []containers.ContainerType {
		var c []containers.ContainerType
		if !useLocalContainers {
			c = append(c, containers.Cassandra)
		}
		return c
	}()...)
	if err != nil {
		log.Error().Msgf("error starting containers: %v", err)
		return
	}
	time.Sleep(10 * time.Second)
	cassandraPort := 9042
	if !useLocalContainers {
		cassandraPort = *c[containers.Cassandra].RESTPort
	}
	Config.Cassandra.Port = cassandraPort
	grpcServer, lis, err := server.GetGRPCServer(Config.Server, &Config.Cassandra)
	if err != nil {
		log.Fatal().Msgf("failed to build grpcServer: %v", err)
	}
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal().Msgf("failed to build grpcServer: %v", err)
		}
	}()
	jaeger.Close()

	ctx.BeforeSuite(func() {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", Config.Server.GrpcPort),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Error().AnErr("err ", err)
		}
		clientFoo := genericsapiv1.NewFooServiceClient(conn)
		clientBar := genericsapiv1.NewBarServiceClient(conn)
		w = NewWorld(clientFoo, clientBar, &Config)
	})

	ctx.AfterSuite(func() {
		dockerEnv.PurgeContainers()
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		w.DeleteTableRecords()
		w.NewData()
		return ctx, nil
	})
	ctx.Step(`^I have a clean (.*) table$`, w.IHaveACleanTable)
	ctx.Step(`^I insert (\d+) value(s|) into (.*) table$`, w.IInsertSomeValues)
	ctx.Step(`^I should have (\d+) entit(y|ies)$`, w.IShouldHaveEntities)
	ctx.Step(`^I should have (\d+) page(s|)$`, w.IShouldHavePages)
	ctx.Step(`^I add the field (.*) with value (.*) to filter$`, w.IAddAFilter)
	ctx.Step(`^I add the current (Timestamp|Day) to filter$`, w.IAddATimestampFilter)
	ctx.Step(`^I query the (.*) service(| with filter)$`, w.IQueryTheService)
	ctx.Step(`^I query the (.*) service with pages of size (\d+)$`, w.IQueryTheServiceWithPages)
}
