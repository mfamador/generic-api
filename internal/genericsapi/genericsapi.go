// Package genericsapi has the implementation of generics API services
package genericsapi

import (
	"genericsapi/internal/genericsapiv1"
	"genericsapi/internal/model"
	"math"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const precision = 5

var n10 = math.Pow10(precision)

func round(f float64) float64 {
	var z float64
	if f < 0 {
		z = -1
	}
	return math.Trunc((f+z*0.5/n10)*n10) / n10
}

func mapFoo(cs *string, res []*model.Foo) (*genericsapiv1.ReadFooReply, error) {
	return &genericsapiv1.ReadFooReply{
		Cursor: func() string {
			if cs != nil {
				return *cs
			}
			return ""
		}(),
		Foos: func() []*genericsapiv1.Foo {
			var foos = make([]*genericsapiv1.Foo, len(res))
			for i := range res {
				vr := res[i]
				foos[i] = &genericsapiv1.Foo{
					Id:          vr.ID,
					Name:        vr.Name,
					Value:       vr.Value,
					FooSpecific: vr.SpecificFoo,
					Timestamp:   timestamppb.New(vr.Timestamp),
				}
			}
			return foos
		}(),
	}, nil
}

func mapBar(cs *string, res []*model.Bar) (*genericsapiv1.ReadBarReply, error) {
	return &genericsapiv1.ReadBarReply{
		Cursor: func() string {
			if cs != nil {
				return *cs
			}
			return ""
		}(),
		Bars: func() []*genericsapiv1.Bar {
			var bars = make([]*genericsapiv1.Bar, len(res))
			for i := range res {
				vr := res[i]
				bars[i] = &genericsapiv1.Bar{
					Id:          vr.ID,
					Name:        vr.Name,
					Value:       vr.Value,
					BarSpecific: vr.SpecificBar,
					Timestamp:   timestamppb.New(vr.Timestamp),
				}
			}
			return bars
		}(),
	}, nil
}
