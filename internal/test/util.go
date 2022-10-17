// Package test for the application
package test

import (
	"errors"
	"genericsapi/internal/config"
	"genericsapi/internal/genericsapiv1"
	"time"
)

// ExpectThen returns the next func
func ExpectThen(msg string, next func() error) error {
	if msg != "" {
		return errors.New(msg)
	}
	return next()
}

// Expect returns the next func
func Expect(msg string) error {
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}

type filter struct {
	key   string
	value string
}

// ScenarioData holds the data for a scenario run
type ScenarioData struct {
	Foos    []*genericsapiv1.Foo
	Bars    []*genericsapiv1.Bar
	Service string
	Pages   int
	Count   int
	Filter  []filter
}

// World holds global vars to be used in tests
type World struct {
	TestStart int64
	Config    *config.AppConfig
	Data      *ScenarioData
	ClientFoo genericsapiv1.FooServiceClient
	ClientBar genericsapiv1.BarServiceClient
	tables    []string
}

// NewWorld creates a new World structure
func NewWorld(clientFoo genericsapiv1.FooServiceClient, clientBar genericsapiv1.BarServiceClient, cnfg *config.AppConfig) *World {
	return &World{
		TestStart: time.Now().Unix(),
		Data:      &ScenarioData{},
		ClientFoo: clientFoo,
		ClientBar: clientBar,
		Config:    cnfg,
		tables: []string{
			"foo", "bar",
		},
	}
}

// DeleteCassandraTablesTables deletes all the cassandra tables
func (w *World) DeleteCassandraTablesTables() {

	// TODO
}

// DeleteTableRecords deletes all the records
func (w *World) DeleteTableRecords() {
	for range w.tables {
		// TODO
	}
}

// NewData creates a new clean scenario data
func (w *World) NewData() {
	w.Data = &ScenarioData{
		Foos:    nil,
		Bars:    nil,
		Service: "",
		Pages:   0,
		Count:   0,
		Filter:  []filter{},
	}
}
