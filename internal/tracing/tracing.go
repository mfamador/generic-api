// Package tracing setups the tracing client
package tracing

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// Config defines the handler configuration
type Config struct {
	AgentAddress string `yaml:"agentAddress"`
}

// Jaeger is a tracer with the capability to push spans to a Jaeger instance.
type Jaeger struct {
	closer io.Closer
}

// NewJaeger initializes the tracing client
func NewJaeger(conf Config) *Jaeger {
	j := &Jaeger{}

	cfg := jaegercfg.Configuration{
		ServiceName: "data-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	reporterConf := &jaegercfg.ReporterConfig{}
	if i := conf.AgentAddress; len(i) > 0 {
		reporterConf.LocalAgentHostPort = i
		cfg.Reporter = reporterConf
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	j.closer = closer
	opentracing.SetGlobalTracer(tracer)

	return j
}

// Close stops the tracer.
func (j *Jaeger) Close() error {
	if j.closer != nil {
		j.closer.Close()
		j.closer = nil
	}
	return nil
}
