// Package service contains the business logic for services
package service

import (
	"errors"
	"fmt"
	"genericsapi/internal/genericsapiv1"
	"reflect"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/rs/zerolog/log"
)

func hasCursor(cursor *string) bool {
	return func(cursor *string) bool { return cursor != nil }(cursor)
}

func logRequest(t, s string, cursor *string, filter []*genericsapiv1.Filter) {
	log.Debug().
		Interface("filter", filter).
		Bool("cursor", hasCursor(cursor)).
		Msg(fmt.Sprintf("%s request traceID=%s", s, t))
}

func logResponse(t, s string, cursor *string, err error) {
	log.Debug().
		Bool("cursor", hasCursor(cursor)).
		AnErr("err", err).
		Msg(fmt.Sprintf("%s response traceID=%s", s, t))
}

func startSpan(s interface{}) (nm string, sp opentracing.Span, tid string) {
	sn := reflect.TypeOf(s).Elem().Name()
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(sn)
	traceID := getSpanID(span)
	return sn, span, traceID
}

func getSpanID(span opentracing.Span) string {
	traceID := ""
	switch sctx := span.Context().(type) {
	case jaeger.SpanContext:
		traceID = sctx.TraceID().String()
	default:
		log.Debug().Msgf("%v", span.Context())
	}
	return traceID
}

func validateFilter(filters []*genericsapiv1.Filter) error {
	for _, filter := range filters {
		if filter.ColumnName == "" {
			return errors.New("column name cannot be null")
		}
	}
	return nil
}
