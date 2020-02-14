package lightstep_tracer_common_test

import (
	"reflect"
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb/collectorpbfakes"
	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/lightsteppb"
	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/metricspb"
)

func TestProtoIsProtobuf(t *testing.T) {
	var r collectorpb.ReportRequest
	if ptag := reflect.ValueOf(r).Type().Field(0).Tag.Get("protobuf"); ptag == "" {
		panic("Not a protobuf!")
	}

	var c lightsteppb.BinaryCarrier
	if ptag := reflect.ValueOf(c).Type().Field(0).Tag.Get("protobuf"); ptag == "" {
		panic("Not a protobuf!")
	}

	var m metricspb.IngestRequest
	if ptag := reflect.ValueOf(m).Type().Field(0).Tag.Get("protobuf"); ptag == "" {
		panic("Not a protobuf!")
	}

	_ = collectorpbfakes.FakeCollectorServiceClient{}
}
