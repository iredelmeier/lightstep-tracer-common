// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/lightstep.proto

/*
Package lightsteppb is a generated protocol buffer package.

It is generated from these files:
	github.com/lightstep/lightstep-tracer-common/lightstep.proto

It has these top-level messages:
	BinaryCarrier
	BasicTracerCarrier
*/
package lightsteppb // import "github.com/lightstep/lightstep-tracer-common/golang/protobuf/lightsteppb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The standard carrier for binary context propagation into LightStep.
type BinaryCarrier struct {
	// "text_ctx" was deprecated following lightstep-tracer-cpp-0.36
	DeprecatedTextCtx [][]byte `protobuf:"bytes,1,rep,name=deprecated_text_ctx,json=deprecatedTextCtx,proto3" json:"deprecated_text_ctx,omitempty"`
	// The Opentracing "basictracer" proto.
	BasicCtx *BasicTracerCarrier `protobuf:"bytes,2,opt,name=basic_ctx,json=basicCtx" json:"basic_ctx,omitempty"`
}

func (m *BinaryCarrier) Reset()                    { *m = BinaryCarrier{} }
func (m *BinaryCarrier) String() string            { return proto.CompactTextString(m) }
func (*BinaryCarrier) ProtoMessage()               {}
func (*BinaryCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BinaryCarrier) GetDeprecatedTextCtx() [][]byte {
	if m != nil {
		return m.DeprecatedTextCtx
	}
	return nil
}

func (m *BinaryCarrier) GetBasicCtx() *BasicTracerCarrier {
	if m != nil {
		return m.BasicCtx
	}
	return nil
}

// Copy of https://github.com/opentracing/basictracer-go/blob/master/wire/wire.proto
type BasicTracerCarrier struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BasicTracerCarrier) Reset()                    { *m = BasicTracerCarrier{} }
func (m *BasicTracerCarrier) String() string            { return proto.CompactTextString(m) }
func (*BasicTracerCarrier) ProtoMessage()               {}
func (*BasicTracerCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BasicTracerCarrier) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *BasicTracerCarrier) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*BinaryCarrier)(nil), "lightstep.BinaryCarrier")
	proto.RegisterType((*BasicTracerCarrier)(nil), "lightstep.BasicTracerCarrier")
}

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/lightstep.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0xc7, 0xd9, 0x46, 0xdb, 0x66, 0xda, 0x82, 0x5d, 0x05, 0xa3, 0x20, 0x84, 0x9e, 0x72, 0x69,
	0x0a, 0xf5, 0x22, 0x45, 0x10, 0x52, 0x3c, 0xf4, 0xba, 0xf4, 0xe4, 0x25, 0x6c, 0x92, 0x21, 0x5d,
	0x6c, 0x1e, 0x6c, 0xa6, 0x92, 0xe0, 0x27, 0xf7, 0x26, 0xd9, 0x6a, 0x23, 0x14, 0xbc, 0xcd, 0xff,
	0x31, 0xcc, 0x0f, 0x06, 0x9e, 0x53, 0x45, 0xbb, 0x43, 0xe4, 0xc7, 0x45, 0xb6, 0xd8, 0xab, 0x74,
	0x47, 0x15, 0x61, 0xd9, 0x4d, 0x73, 0xd2, 0x32, 0x46, 0x3d, 0x8f, 0x8b, 0x2c, 0x2b, 0xf2, 0xce,
	0xf7, 0x4b, 0x5d, 0x50, 0xc1, 0xed, 0x93, 0x31, 0xfb, 0x84, 0x49, 0xa0, 0x72, 0xa9, 0x9b, 0xb5,
	0xd4, 0x5a, 0xa1, 0xe6, 0x3e, 0x5c, 0x27, 0x58, 0x6a, 0x8c, 0x25, 0x61, 0x12, 0x12, 0xd6, 0x14,
	0xc6, 0x54, 0x3b, 0xcc, 0xb5, 0xbc, 0xb1, 0x98, 0x76, 0xd1, 0x16, 0x6b, 0x5a, 0x53, 0xcd, 0x57,
	0x60, 0x47, 0xb2, 0x52, 0xb1, 0x69, 0xf5, 0x5c, 0xe6, 0x8d, 0x96, 0x0f, 0x7e, 0x77, 0x30, 0x68,
	0xb3, 0xad, 0xa1, 0xf9, 0xb9, 0x20, 0x86, 0xa6, 0xbf, 0xa6, 0x7a, 0xf6, 0xc5, 0x80, 0x9f, 0x17,
	0xf8, 0x1d, 0x0c, 0x0d, 0x7f, 0xa8, 0x12, 0x87, 0xb9, 0xcc, 0xeb, 0x8b, 0x81, 0xd1, 0x9b, 0x84,
	0xdf, 0xc2, 0xa0, 0x2a, 0x65, 0xde, 0x26, 0x3d, 0x93, 0xf4, 0x5b, 0xb9, 0x49, 0xb8, 0x03, 0x83,
	0x4a, 0x66, 0xe5, 0x1e, 0x13, 0xc7, 0x72, 0x99, 0x37, 0x14, 0xbf, 0x92, 0x6f, 0x61, 0x12, 0xc9,
	0x34, 0x95, 0x29, 0x86, 0x8a, 0x30, 0xab, 0x9c, 0x0b, 0xd7, 0xf2, 0x46, 0xcb, 0xc5, 0xbf, 0x90,
	0x7e, 0x70, 0x5c, 0xd9, 0xb4, 0x1b, 0xaf, 0x39, 0xe9, 0x46, 0x8c, 0xa3, 0x3f, 0xd6, 0xfd, 0x0b,
	0x4c, 0xcf, 0x2a, 0xfc, 0x0a, 0xac, 0x77, 0x6c, 0x0c, 0xb3, 0x2d, 0xda, 0x91, 0xdf, 0xc0, 0xe5,
	0x87, 0xdc, 0x1f, 0xd0, 0xd0, 0xda, 0xe2, 0x28, 0x56, 0xbd, 0x27, 0x16, 0x4c, 0xde, 0x46, 0x27,
	0x80, 0x32, 0x8a, 0xfa, 0xe6, 0x33, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x78, 0x62,
	0xbc, 0xd9, 0x01, 0x00, 0x00,
}
