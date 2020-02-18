// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/tmpgen/metrics.proto

/*
	Package metricspb is a generated protocol buffer package.

	It is generated from these files:
		github.com/lightstep/lightstep-tracer-common/tmpgen/metrics.proto

	It has these top-level messages:
		MetricPoint
		IngestRequest
		IngestResponse
*/
package metricspb // import "github.com/lightstep/lightstep-tracer-common/golang/gogo/metricspb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Kind indicates the semantics of the points.
type MetricKind int32

const (
	MetricKind_INVALID_METRIC_KIND MetricKind = 0
	MetricKind_DELTA               MetricKind = 1
	MetricKind_GAUGE               MetricKind = 2
)

var MetricKind_name = map[int32]string{
	0: "INVALID_METRIC_KIND",
	1: "DELTA",
	2: "GAUGE",
}
var MetricKind_value = map[string]int32{
	"INVALID_METRIC_KIND": 0,
	"DELTA":               1,
	"GAUGE":               2,
}

func (x MetricKind) String() string {
	return proto.EnumName(MetricKind_name, int32(x))
}
func (MetricKind) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{0} }

type MetricPoint struct {
	Kind          MetricKind        `protobuf:"varint,1,opt,name=kind,proto3,enum=lightstep.metrics.MetricKind" json:"kind,omitempty"`
	TimeSeriesKey string            `protobuf:"bytes,2,opt,name=time_series_key,json=timeSeriesKey,proto3" json:"time_series_key,omitempty"`
	Labels        map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to Value:
	//	*MetricPoint_Uint
	//	*MetricPoint_Float
	Value isMetricPoint_Value `protobuf_oneof:"value"`
}

func (m *MetricPoint) Reset()                    { *m = MetricPoint{} }
func (m *MetricPoint) String() string            { return proto.CompactTextString(m) }
func (*MetricPoint) ProtoMessage()               {}
func (*MetricPoint) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{0} }

type isMetricPoint_Value interface {
	isMetricPoint_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MetricPoint_Uint struct {
	Uint uint64 `protobuf:"varint,4,opt,name=uint,proto3,oneof"`
}
type MetricPoint_Float struct {
	Float float64 `protobuf:"fixed64,5,opt,name=float,proto3,oneof"`
}

func (*MetricPoint_Uint) isMetricPoint_Value()  {}
func (*MetricPoint_Float) isMetricPoint_Value() {}

func (m *MetricPoint) GetValue() isMetricPoint_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricPoint) GetKind() MetricKind {
	if m != nil {
		return m.Kind
	}
	return MetricKind_INVALID_METRIC_KIND
}

func (m *MetricPoint) GetTimeSeriesKey() string {
	if m != nil {
		return m.TimeSeriesKey
	}
	return ""
}

func (m *MetricPoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricPoint) GetUint() uint64 {
	if x, ok := m.GetValue().(*MetricPoint_Uint); ok {
		return x.Uint
	}
	return 0
}

func (m *MetricPoint) GetFloat() float64 {
	if x, ok := m.GetValue().(*MetricPoint_Float); ok {
		return x.Float
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricPoint) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricPoint_OneofMarshaler, _MetricPoint_OneofUnmarshaler, _MetricPoint_OneofSizer, []interface{}{
		(*MetricPoint_Uint)(nil),
		(*MetricPoint_Float)(nil),
	}
}

func _MetricPoint_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricPoint)
	// value
	switch x := m.Value.(type) {
	case *MetricPoint_Uint:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Uint))
	case *MetricPoint_Float:
		_ = b.EncodeVarint(5<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Float))
	case nil:
	default:
		return fmt.Errorf("MetricPoint.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricPoint_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricPoint)
	switch tag {
	case 4: // value.uint
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricPoint_Uint{x}
		return true, err
	case 5: // value.float
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricPoint_Float{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _MetricPoint_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricPoint)
	// value
	switch x := m.Value.(type) {
	case *MetricPoint_Uint:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint))
	case *MetricPoint_Float:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type IngestRequest struct {
	IdempotencyKey string                      `protobuf:"bytes,1,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	ReporterId     uint64                      `protobuf:"varint,2,opt,name=reporter_id,json=reporterId,proto3" json:"reporter_id,omitempty"`
	Labels         map[string]string           `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Start          *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=start" json:"start,omitempty"`
	Duration       *google_protobuf.Duration   `protobuf:"bytes,5,opt,name=duration" json:"duration,omitempty"`
	Points         []*MetricPoint              `protobuf:"bytes,6,rep,name=points" json:"points,omitempty"`
}

func (m *IngestRequest) Reset()                    { *m = IngestRequest{} }
func (m *IngestRequest) String() string            { return proto.CompactTextString(m) }
func (*IngestRequest) ProtoMessage()               {}
func (*IngestRequest) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{1} }

func (m *IngestRequest) GetIdempotencyKey() string {
	if m != nil {
		return m.IdempotencyKey
	}
	return ""
}

func (m *IngestRequest) GetReporterId() uint64 {
	if m != nil {
		return m.ReporterId
	}
	return 0
}

func (m *IngestRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *IngestRequest) GetStart() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IngestRequest) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *IngestRequest) GetPoints() []*MetricPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

type IngestResponse struct {
}

func (m *IngestResponse) Reset()                    { *m = IngestResponse{} }
func (m *IngestResponse) String() string            { return proto.CompactTextString(m) }
func (*IngestResponse) ProtoMessage()               {}
func (*IngestResponse) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{2} }

func init() {
	proto.RegisterType((*MetricPoint)(nil), "lightstep.metrics.MetricPoint")
	proto.RegisterType((*IngestRequest)(nil), "lightstep.metrics.IngestRequest")
	proto.RegisterType((*IngestResponse)(nil), "lightstep.metrics.IngestResponse")
	proto.RegisterEnum("lightstep.metrics.MetricKind", MetricKind_name, MetricKind_value)
}
func (m *MetricPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricPoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Kind))
	}
	if len(m.TimeSeriesKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.TimeSeriesKey)))
		i += copy(dAtA[i:], m.TimeSeriesKey)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovMetrics(uint64(len(k))) + 1 + len(v) + sovMetrics(uint64(len(v)))
			i = encodeVarintMetrics(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetrics(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetrics(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.Value != nil {
		nn1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *MetricPoint_Uint) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x20
	i++
	i = encodeVarintMetrics(dAtA, i, uint64(m.Uint))
	return i, nil
}
func (m *MetricPoint_Float) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x29
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Float))))
	i += 8
	return i, nil
}
func (m *IngestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngestRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IdempotencyKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.IdempotencyKey)))
		i += copy(dAtA[i:], m.IdempotencyKey)
	}
	if m.ReporterId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.ReporterId))
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovMetrics(uint64(len(k))) + 1 + len(v) + sovMetrics(uint64(len(v)))
			i = encodeVarintMetrics(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetrics(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetrics(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.Start != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Start.Size()))
		n2, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Duration != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Duration.Size()))
		n3, err := m.Duration.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Points) > 0 {
		for _, msg := range m.Points {
			dAtA[i] = 0x32
			i++
			i = encodeVarintMetrics(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *IngestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngestResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MetricPoint) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovMetrics(uint64(m.Kind))
	}
	l = len(m.TimeSeriesKey)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMetrics(uint64(len(k))) + 1 + len(v) + sovMetrics(uint64(len(v)))
			n += mapEntrySize + 1 + sovMetrics(uint64(mapEntrySize))
		}
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *MetricPoint_Uint) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetrics(uint64(m.Uint))
	return n
}
func (m *MetricPoint_Float) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *IngestRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.IdempotencyKey)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.ReporterId != 0 {
		n += 1 + sovMetrics(uint64(m.ReporterId))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMetrics(uint64(len(k))) + 1 + len(v) + sovMetrics(uint64(len(v)))
			n += mapEntrySize + 1 + sovMetrics(uint64(mapEntrySize))
		}
	}
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Duration != nil {
		l = m.Duration.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	return n
}

func (m *IngestResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovMetrics(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (MetricKind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeSeriesKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeSeriesKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMetrics
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMetrics
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMetrics
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthMetrics
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMetrics(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMetrics
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &MetricPoint_Uint{v}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Float", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = &MetricPoint_Float{float64(math.Float64frombits(v))}
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IngestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IngestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdempotencyKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdempotencyKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReporterId", wireType)
			}
			m.ReporterId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReporterId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMetrics
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMetrics
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMetrics
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthMetrics
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMetrics(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMetrics
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &google_protobuf1.Timestamp{}
			}
			if err := m.Start.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &google_protobuf.Duration{}
			}
			if err := m.Duration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, &MetricPoint{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IngestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IngestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetrics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetrics(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/tmpgen/metrics.proto", fileDescriptorMetrics)
}

var fileDescriptorMetrics = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdb, 0x6e, 0xd3, 0x30,
	0x18, 0xae, 0xb3, 0xb4, 0x50, 0x47, 0xeb, 0x8a, 0x99, 0x20, 0xab, 0x44, 0xa8, 0x8a, 0x04, 0xd5,
	0xc4, 0xd2, 0x51, 0x04, 0xe2, 0x70, 0x81, 0x5a, 0x5a, 0x6d, 0x51, 0xbb, 0xa9, 0xca, 0x0a, 0x17,
	0xdc, 0x44, 0x69, 0xe2, 0x65, 0xd6, 0x92, 0x38, 0xd8, 0x0e, 0x52, 0xdf, 0x82, 0x27, 0xe0, 0x82,
	0xa7, 0xe1, 0x92, 0x47, 0x40, 0xe5, 0x41, 0x40, 0x71, 0xd2, 0xb5, 0x6c, 0x83, 0x1b, 0xee, 0xfe,
	0xc3, 0xf7, 0x9f, 0xbe, 0xcf, 0x86, 0xbd, 0x80, 0x88, 0xb3, 0x74, 0x66, 0x7a, 0x34, 0xea, 0x84,
	0x24, 0x38, 0x13, 0x5c, 0xe0, 0x64, 0x65, 0xed, 0x09, 0xe6, 0x7a, 0x98, 0xed, 0x79, 0x34, 0x8a,
	0x68, 0xdc, 0x11, 0x51, 0x12, 0xe0, 0xb8, 0x13, 0x61, 0xc1, 0x88, 0xc7, 0xcd, 0x84, 0x51, 0x41,
	0xd1, 0xad, 0x0b, 0xb4, 0x59, 0x24, 0x1a, 0x46, 0x40, 0x69, 0x10, 0xe2, 0x8e, 0x04, 0xcc, 0xd2,
	0xd3, 0x8e, 0x9f, 0x32, 0x57, 0x10, 0x1a, 0xe7, 0x25, 0x8d, 0xfb, 0x97, 0xf3, 0x82, 0x44, 0x98,
	0x0b, 0x37, 0x4a, 0x72, 0x40, 0xeb, 0x8b, 0x02, 0xb5, 0x23, 0xd9, 0x6c, 0x42, 0x49, 0x2c, 0xd0,
	0x13, 0xa8, 0x9e, 0x93, 0xd8, 0xd7, 0x41, 0x13, 0xb4, 0x6b, 0xdd, 0x7b, 0xe6, 0x95, 0x91, 0x66,
	0x8e, 0x1e, 0x91, 0xd8, 0xb7, 0x25, 0x14, 0x3d, 0x84, 0x5b, 0x59, 0x57, 0x87, 0x63, 0x46, 0x30,
	0x77, 0xce, 0xf1, 0x5c, 0x57, 0x9a, 0xa0, 0x5d, 0xb5, 0x37, 0xb3, 0xf0, 0x89, 0x8c, 0x8e, 0xf0,
	0x1c, 0xf5, 0x61, 0x25, 0x74, 0x67, 0x38, 0xe4, 0xfa, 0x46, 0x73, 0xa3, 0xad, 0x75, 0x77, 0xff,
	0xda, 0x5c, 0xae, 0x62, 0x8e, 0x25, 0x78, 0x18, 0x0b, 0x36, 0xb7, 0x8b, 0x4a, 0xb4, 0x0d, 0xd5,
	0x94, 0xc4, 0x42, 0x57, 0x9b, 0xa0, 0xad, 0x1e, 0x96, 0x6c, 0xe9, 0xa1, 0x3b, 0xb0, 0x7c, 0x1a,
	0x52, 0x57, 0xe8, 0xe5, 0x26, 0x68, 0x83, 0xc3, 0x92, 0x9d, 0xbb, 0x8d, 0x97, 0x50, 0x5b, 0x6b,
	0x82, 0xea, 0x70, 0x23, 0x5b, 0x0e, 0xc8, 0xe5, 0x32, 0x13, 0x6d, 0xc3, 0xf2, 0x27, 0x37, 0x4c,
	0x71, 0xb1, 0x70, 0xee, 0xbc, 0x52, 0x5e, 0x80, 0xfe, 0x8d, 0x22, 0xd3, 0xfa, 0xa5, 0xc0, 0x4d,
	0x2b, 0x0e, 0x30, 0x17, 0x36, 0xfe, 0x98, 0x62, 0x2e, 0xd0, 0x23, 0xb8, 0x45, 0x7c, 0x1c, 0x25,
	0x54, 0xe0, 0xd8, 0x9b, 0x3b, 0xab, 0x96, 0xb5, 0xb5, 0x70, 0x76, 0xf0, 0x03, 0xa8, 0x31, 0x9c,
	0x50, 0x26, 0x30, 0x73, 0x88, 0x2f, 0x67, 0xa8, 0x7d, 0x65, 0x1f, 0xd8, 0x70, 0x19, 0xb6, 0x7c,
	0x34, 0xb8, 0xc4, 0xca, 0xe3, 0x6b, 0x58, 0xf9, 0x63, 0xfe, 0xb5, 0xbc, 0xec, 0xc3, 0x32, 0x17,
	0x2e, 0xcb, 0x89, 0xd1, 0xba, 0x0d, 0x33, 0xd7, 0xdd, 0x5c, 0xea, 0x6e, 0x4e, 0x97, 0xba, 0xdb,
	0x39, 0x10, 0x3d, 0x83, 0x37, 0x97, 0x6f, 0x45, 0xd2, 0xa6, 0x75, 0x77, 0xae, 0x14, 0x0d, 0x0a,
	0x80, 0x7d, 0x01, 0x45, 0xcf, 0x61, 0x25, 0xc9, 0xd4, 0xe1, 0x7a, 0x45, 0xae, 0x6b, 0xfc, 0x5b,
	0x44, 0xbb, 0x40, 0xff, 0x87, 0x14, 0xad, 0x3a, 0xac, 0x2d, 0x09, 0xe0, 0x09, 0x8d, 0x39, 0xde,
	0x7d, 0x0d, 0xe1, 0xea, 0x15, 0xa2, 0xbb, 0xf0, 0xb6, 0x75, 0xfc, 0xbe, 0x37, 0xb6, 0x06, 0xce,
	0xd1, 0x70, 0x6a, 0x5b, 0x6f, 0x9d, 0x91, 0x75, 0x3c, 0xa8, 0x97, 0x50, 0x15, 0x96, 0x07, 0xc3,
	0xf1, 0xb4, 0x57, 0x07, 0x99, 0x79, 0xd0, 0x7b, 0x77, 0x30, 0xac, 0x2b, 0xfd, 0x37, 0xdf, 0x16,
	0x06, 0xf8, 0xbe, 0x30, 0xc0, 0x8f, 0x85, 0x01, 0x3e, 0xff, 0x34, 0x4a, 0x70, 0xc7, 0xa3, 0xd1,
	0xda, 0x19, 0xf9, 0x4f, 0x34, 0x03, 0x96, 0x78, 0x13, 0xf0, 0xa1, 0x5a, 0x5c, 0x95, 0xcc, 0xbe,
	0x2a, 0xea, 0xf8, 0x64, 0xd2, 0x9f, 0x55, 0x24, 0x3f, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0xf8, 0x69, 0x4e, 0xd2, 0x03, 0x00, 0x00,
}
