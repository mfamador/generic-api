// Package test for the application
package test

import (
	"context"
	"errors"
	"genericsapi/internal/repository"
	"strconv"
	"strings"
	"time"

	"github.com/smartystreets/assertions"
)

const (
	readings                         = "Readings"
	readingsLastValues               = "ReadingsLastValues"
	readingsLastValuesManual         = "ReadingsLastValuesManual"
	volumeReadings                   = "VolumeReadings"
	volumeReadingsLastValues         = "VolumeReadingsLastValues"
	volumeReadingsLastValuesManual   = "VolumeReadingsLastValuesManual"
	consumptions                     = "Consumptions"
	consumptionsLastValues           = "ConsumptionsLastValues"
	refills                          = "Refills"
	refillsLastValues                = "RefillsLastValues"
	batteryData                      = "BatteryData"
	batteryDataLastValues            = "BatteryDataLastValues"
	batteryDataLastValuesManual      = "BatteryDataLastValuesManual"
	networkQuality                   = "NetworkQuality"
	networkQualityLastValues         = "NetworkQualityLastValues"
	networkQualityLastValuesManual   = "NetworkQualityLastValuesManual"
	packets                          = "Packets"
	packetsLastValues                = "PacketsLastValues"
	events                           = "Events"
	eventsLastValues                 = "EventsLastValues"
	timeout                          = 10
	unknownServiceErrorMsg           = "unknown service"
	errorOccurredErrorMsg            = "an error has occurred"
	shouldHaveReturnedValuesErrorMsg = "should have returned values"
	imei                             = "imei"
	serialNumber                     = "serial_number"
	deviceTypeVariant                = "deviceTypeVariant"
)

// IHaveACleanTable deletes the table
func (w *World) IHaveACleanTable(tableName string) error {
	const timeout = 30
	// TODO
	return nil
}

// IInsertSomeValues stores sample values
func (w *World) IInsertSomeValues(n int, _, tableName string) error { //nolint
	// TODO
	return nil
}

// IAddATimestampFilter stores the current timestamp
func (w *World) IAddATimestampFilter(filterName string) error {
	return w.IAddAFilter(filterName, "")
}

// IAddAFilter stores a filter
func (w *World) IAddAFilter(filterName, value string) error {
	switch filterName {
	case repository.Day, repository.Timestamp:
		time.Sleep(2 * time.Second) //nolint:gomnd
		w.Data.Filter = append(w.Data.Filter, filter{
			key:   filterName,
			value: strconv.Itoa(int(time.Now().Unix())),
		})
	default:
		w.Data.Filter = append(w.Data.Filter, filter{
			key:   filterName,
			value: strings.TrimSpace(value),
		})
	}
	return nil
}

// IQueryTheServiceWithPages queries last values table
//
//nolint:funlen,gocyclo
func (w *World) IQueryTheServiceWithPages(service string, limit int) error {
	filter := w.getFilter()

	w.Data.Service = service
	switch service {
	case readings:
		err := w.queryReadings(limit, filter)
		if err != nil {
			return err
		}
	case readingsLastValues, readingsLastValuesManual:
		err := w.queryReadingsLastValues(limit, filter)
		if err != nil {
			return err
		}
	case volumeReadings:
		err := w.queryVolumeReadings(limit, filter)
		if err != nil {
			return err
		}
	case volumeReadingsLastValues, volumeReadingsLastValuesManual:
		err := w.queryVolumeReadingsLastValues(limit, filter)
		if err != nil {
			return err
		}
	case consumptions:
		err := w.queryConsumptions(limit, filter)
		if err != nil {
			return err
		}
	case consumptionsLastValues:
		err := w.queryConsumptionsLastValues(limit, filter)
		if err != nil {
			return err
		}
	case refills:
		err := w.queryRefills(limit, filter)
		if err != nil {
			return err
		}
	case refillsLastValues:
		err := w.queryRefillsLastValues(limit, filter)
		if err != nil {
			return err
		}
	case batteryData:
		err := w.queryBatteryData(limit, filter)
		if err != nil {
			return err
		}
	case batteryDataLastValues, batteryDataLastValuesManual:
		err := w.queryBatteryDataLastValues(limit, filter)
		if err != nil {
			return err
		}
	case networkQuality:
		err := w.queryNetworkQuality(limit, filter)
		if err != nil {
			return err
		}
	case networkQualityLastValues, networkQualityLastValuesManual:
		err := w.queryNetworkQualityLastValues(limit, filter)
		if err != nil {
			return err
		}
	case packets:
		err := w.queryPackets(limit, filter)
		if err != nil {
			return err
		}
	case packetsLastValues:
		err := w.queryPacketsLastValues(limit, filter)
		if err != nil {
			return err
		}
	case events:
		err := w.queryEvents(limit, filter)
		if err != nil {
			return err
		}
	case eventsLastValues:
		err := w.queryEventsLastValues(limit, filter)
		if err != nil {
			return err
		}
	default:
		return errors.New(unknownServiceErrorMsg)
	}

	return nil
}

// IQueryTheService queries a service
func (w *World) IQueryTheService(service string) error {
	return w.IQueryTheServiceWithPages(service, int(w.Config.DB.Limit))
}

func (w *World) queryNetworkQuality(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadNetworkQualityReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadNetworkQuality(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.NetworkQuality
		w.Data.NetworkQuality = append(w.Data.NetworkQuality, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryBatteryData(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadBatteryDataReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadBatteryData(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.BatteryData
		w.Data.BatteryData = append(w.Data.BatteryData, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryVolumeReadings(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadVolumeReadingsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadVolumeReadings(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.VolumeReadings
		w.Data.VolumeReadings = append(w.Data.VolumeReadings, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryConsumptions(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadConsumptionsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadConsumptions(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Consumptions
		w.Data.Consumptions = append(w.Data.Consumptions, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryRefills(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadRefillsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadRefills(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Refills
		w.Data.Refills = append(w.Data.Refills, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryReadings(limit int, filters []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filters,
	}
	var err error
	var queryReply *dataapiv1.ReadReadingsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadReadings(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Readings
		w.Data.Readings = append(w.Data.Readings, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryReadingsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadReadingsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadReadingsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Readings
		w.Data.ReadingsLastValues = append(w.Data.ReadingsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryVolumeReadingsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadVolumeReadingsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadVolumeReadingsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.VolumeReadings
		w.Data.VolumeReadingsLastValues = append(w.Data.VolumeReadingsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryConsumptionsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadConsumptionsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadConsumptionsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Consumptions
		w.Data.ConsumptionsLastValues = append(w.Data.ConsumptionsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryRefillsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadRefillsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadRefillsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Refills
		w.Data.RefillsLastValues = append(w.Data.RefillsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryBatteryDataLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadBatteryDataReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadBatteryDataLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.BatteryData
		w.Data.BatteryDataLastValues = append(w.Data.BatteryDataLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryNetworkQualityLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadNetworkQualityReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadNetworkQualityLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.NetworkQuality
		w.Data.NetworkQualityLastValues = append(w.Data.NetworkQualityLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryPackets(limit int, filters []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filters,
	}
	var err error
	var queryReply *dataapiv1.ReadPacketsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadPackets(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Packets
		w.Data.Packets = append(w.Data.Packets, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryPacketsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadPacketsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadPacketsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Packets
		w.Data.PacketsLastValues = append(w.Data.PacketsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryEvents(limit int, filters []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filters,
	}
	var err error
	var queryReply *dataapiv1.ReadEventsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadEvents(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Events
		w.Data.Events = append(w.Data.Events, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) queryEventsLastValues(limit int, filter []*dataapiv1.Filter) error { //nolint
	request := dataapiv1.ReadRequest{
		Limit:   uint32(limit),
		Filters: filter,
	}
	var err error
	var queryReply *dataapiv1.ReadEventsReply
	for {
		if queryReply != nil {
			request.Cursor = queryReply.Cursor
		}
		queryReply, err = w.Client.ReadEventsLastValues(context.Background(), &request)
		if err != nil {
			return err
		}
		lv := queryReply.Events
		w.Data.EventsLastValues = append(w.Data.EventsLastValues, lv...)
		if len(lv) > 0 {
			w.Data.Pages++
		}
		if queryReply.Cursor == "" {
			break
		}
	}
	return nil
}

func (w *World) getFilter() []*dataapiv1.Filter {
	var filters []*dataapiv1.Filter

	for _, filter := range w.Data.Filter {
		switch filter.key {
		case repository.Timestamp:
			v, _ := strconv.ParseInt(filter.value, 10, 64) //nolint:gomnd
			ts := time.Unix(v, 0)
			filter := &dataapiv1.Filter{
				ColumnName:  repository.Timestamp,
				Operator:    dataapiv1.FilterOperator_GREATERTHANOREQUAL,
				StringValue: ts.Format(time.RFC3339),
			}
			filters = append(filters, filter)
		case repository.Day:
			v, _ := strconv.ParseInt(filter.value, 10, 64) //nolint:gomnd
			ts := time.Unix(v, 0)
			filter := &dataapiv1.Filter{
				ColumnName:  repository.PartitionKey,
				Operator:    dataapiv1.FilterOperator_GREATERTHANOREQUAL,
				StringValue: ts.Format("2006-01-02"),
			}
			filters = append(filters, filter)
		default:
			filter := &dataapiv1.Filter{
				ColumnName:  filter.key,
				Operator:    dataapiv1.FilterOperator_EQUAL,
				StringValue: filter.value,
			}
			filters = append(filters, filter)
		}
	}

	return filters
}

// IShouldHaveEntities checks if entities are returned
//
//nolint:gocyclo
func (w *World) IShouldHaveEntities(n int) error {
	switch w.Data.Service {
	case readings:
		return Expect(assertions.ShouldEqual(len(w.Data.Readings), n))
	case readingsLastValues, readingsLastValuesManual:
		return Expect(assertions.ShouldEqual(len(w.Data.ReadingsLastValues), n))
	case volumeReadings:
		return Expect(assertions.ShouldEqual(len(w.Data.VolumeReadings), n))
	case volumeReadingsLastValues, volumeReadingsLastValuesManual:
		return Expect(assertions.ShouldEqual(len(w.Data.VolumeReadingsLastValues), n))
	case consumptions:
		return Expect(assertions.ShouldEqual(len(w.Data.Consumptions), n))
	case consumptionsLastValues:
		return Expect(assertions.ShouldEqual(len(w.Data.ConsumptionsLastValues), n))
	case refills:
		return Expect(assertions.ShouldEqual(len(w.Data.Refills), n))
	case refillsLastValues:
		return Expect(assertions.ShouldEqual(len(w.Data.RefillsLastValues), n))
	case batteryData:
		return Expect(assertions.ShouldEqual(len(w.Data.BatteryData), n))
	case batteryDataLastValues, batteryDataLastValuesManual:
		return Expect(assertions.ShouldEqual(len(w.Data.BatteryDataLastValues), n))
	case networkQuality:
		return Expect(assertions.ShouldEqual(len(w.Data.NetworkQuality), n))
	case networkQualityLastValues, networkQualityLastValuesManual:
		return Expect(assertions.ShouldEqual(len(w.Data.NetworkQualityLastValues), n))
	case packets:
		return Expect(assertions.ShouldEqual(len(w.Data.Packets), n))
	case packetsLastValues:
		return Expect(assertions.ShouldEqual(len(w.Data.PacketsLastValues), n))
	case events:
		return Expect(assertions.ShouldEqual(len(w.Data.Events), n))
	case eventsLastValues:
		return Expect(assertions.ShouldEqual(len(w.Data.EventsLastValues), n))
	}
	return errors.New(unknownServiceErrorMsg)
}

// IShouldHavePages checks if entities are returned
//
//nolint:gocyclo,funlen
func (w *World) IShouldHavePages(n int) error {
	if n != 0 {
		switch w.Data.Service {
		case readings:
			if len(w.Data.Readings) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case readingsLastValues, readingsLastValuesManual:
			if len(w.Data.ReadingsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case volumeReadings:
			if len(w.Data.VolumeReadings) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case volumeReadingsLastValues, volumeReadingsLastValuesManual:
			if len(w.Data.VolumeReadingsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case consumptions:
			if len(w.Data.Consumptions) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case consumptionsLastValues:
			if len(w.Data.ConsumptionsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case refills:
			if len(w.Data.Refills) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case refillsLastValues:
			if len(w.Data.RefillsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case batteryData:
			if len(w.Data.BatteryData) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case batteryDataLastValues, batteryDataLastValuesManual:
			if len(w.Data.BatteryDataLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case networkQuality:
			if len(w.Data.NetworkQuality) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case networkQualityLastValues, networkQualityLastValuesManual:
			if len(w.Data.NetworkQualityLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case packets:
			if len(w.Data.Packets) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case packetsLastValues:
			if len(w.Data.PacketsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case events:
			if len(w.Data.Events) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		case eventsLastValues:
			if len(w.Data.EventsLastValues) == 0 {
				return errors.New(shouldHaveReturnedValuesErrorMsg)
			}
		default:
			return errors.New(unknownServiceErrorMsg)
		}
	}
	return Expect(assertions.ShouldEqual(w.Data.Pages, n))
}
