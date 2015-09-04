// Code generated by protoc-gen-gogo.
// source: cockroach/proto/internal.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// InternalValueType defines a set of string constants placed in the
// "tag" field of Value messages which are created internally. These
// are defined as a protocol buffer enumeration so that they can be
// used portably between our Go and C code. The tags are used by the
// RocksDB Merge Operator to perform specialized merges.
type InternalValueType int32

const (
	// _CR_TS is applied to values which contain InternalTimeSeriesData.
	_CR_TS InternalValueType = 1
)

var InternalValueType_name = map[int32]string{
	1: "_CR_TS",
}
var InternalValueType_value = map[string]int32{
	"_CR_TS": 1,
}

func (x InternalValueType) Enum() *InternalValueType {
	p := new(InternalValueType)
	*p = x
	return p
}
func (x InternalValueType) String() string {
	return proto1.EnumName(InternalValueType_name, int32(x))
}
func (x *InternalValueType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(InternalValueType_value, data, "InternalValueType")
	if err != nil {
		return err
	}
	*x = InternalValueType(value)
	return nil
}

// A RaftCommand is a command which can be serialized and sent via
// raft.
type RaftCommand struct {
	RangeID      RangeID      `protobuf:"varint,1,opt,name=range_id,casttype=RangeID" json:"range_id"`
	OriginNodeID RaftNodeID   `protobuf:"varint,2,opt,name=origin_node_id,casttype=RaftNodeID" json:"origin_node_id"`
	Cmd          BatchRequest `protobuf:"bytes,3,opt,name=cmd" json:"cmd"`
}

func (m *RaftCommand) Reset()         { *m = RaftCommand{} }
func (m *RaftCommand) String() string { return proto1.CompactTextString(m) }
func (*RaftCommand) ProtoMessage()    {}

func (m *RaftCommand) GetRangeID() RangeID {
	if m != nil {
		return m.RangeID
	}
	return 0
}

func (m *RaftCommand) GetOriginNodeID() RaftNodeID {
	if m != nil {
		return m.OriginNodeID
	}
	return 0
}

func (m *RaftCommand) GetCmd() BatchRequest {
	if m != nil {
		return m.Cmd
	}
	return BatchRequest{}
}

// InternalTimeSeriesData is a collection of data samples for some
// measurable value, where each sample is taken over a uniform time
// interval.
//
// The collection itself contains a start timestamp (in seconds since the unix
// epoch) and a sample duration (in milliseconds). Each sample in the collection
// will contain a positive integer offset that indicates the length of time
// between the start_timestamp of the collection and the time when the sample
// began, expressed as an whole number of sample intervals. For example, if the
// sample duration is 60000 (indicating 1 minute), then a contained sample with
// an offset value of 5 begins (5*60000ms = 300000ms = 5 minutes) after the
// start timestamp of this data.
//
// This is meant to be an efficient internal representation of time series data,
// ensuring that very little redundant data is stored on disk. With this goal in
// mind, this message does not identify the variable which is actually being
// measured; that information is expected be encoded in the key where this
// message is stored.
type InternalTimeSeriesData struct {
	// Holds a wall time, expressed as a unix epoch time in nanoseconds. This
	// represents the earliest possible timestamp for a sample within the
	// collection.
	StartTimestampNanos int64 `protobuf:"varint,1,opt,name=start_timestamp_nanos" json:"start_timestamp_nanos"`
	// The duration of each sample interval, expressed in nanoseconds.
	SampleDurationNanos int64 `protobuf:"varint,2,opt,name=sample_duration_nanos" json:"sample_duration_nanos"`
	// The actual data samples for this metric.
	Samples []*InternalTimeSeriesSample `protobuf:"bytes,3,rep,name=samples" json:"samples,omitempty"`
}

func (m *InternalTimeSeriesData) Reset()         { *m = InternalTimeSeriesData{} }
func (m *InternalTimeSeriesData) String() string { return proto1.CompactTextString(m) }
func (*InternalTimeSeriesData) ProtoMessage()    {}

func (m *InternalTimeSeriesData) GetStartTimestampNanos() int64 {
	if m != nil {
		return m.StartTimestampNanos
	}
	return 0
}

func (m *InternalTimeSeriesData) GetSampleDurationNanos() int64 {
	if m != nil {
		return m.SampleDurationNanos
	}
	return 0
}

func (m *InternalTimeSeriesData) GetSamples() []*InternalTimeSeriesSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

// A InternalTimeSeriesSample represents data gathered from multiple
// measurements of a variable value over a given period of time. The
// length of that period of time is stored in an
// InternalTimeSeriesData message; a sample cannot be interpreted
// correctly without a start timestamp and sample duration.
//
// Each sample may contain data gathered from multiple measurements of the same
// variable, as long as all of those measurements occured within the sample
// period. The sample stores several aggregated values from these measurements:
// - The sum of all measured values
// - A count of all measurements taken
// - The maximum individual measurement seen
// - The minimum individual measurement seen
//
// If zero measurements are present in a sample, then it should be omitted
// entirely from any collection it would be a part of.
//
// If the count of measurements is 1, then max and min fields may be omitted
// and assumed equal to the sum field.
type InternalTimeSeriesSample struct {
	// Temporal offset from the "start_timestamp" of the InternalTimeSeriesData
	// collection this data point is part in. The units of this value are
	// determined by the value of the "sample_duration_milliseconds" field of
	// the TimeSeriesData collection.
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// Count of measurements taken within this sample.
	Count uint32 `protobuf:"varint,6,opt,name=count" json:"count"`
	// Sum of all measurements.
	Sum float64 `protobuf:"fixed64,7,opt,name=sum" json:"sum"`
	// Maximum encountered measurement in this sample.
	Max *float64 `protobuf:"fixed64,8,opt,name=max" json:"max,omitempty"`
	// Minimum encountered measurement in this sample.
	Min *float64 `protobuf:"fixed64,9,opt,name=min" json:"min,omitempty"`
}

func (m *InternalTimeSeriesSample) Reset()         { *m = InternalTimeSeriesSample{} }
func (m *InternalTimeSeriesSample) String() string { return proto1.CompactTextString(m) }
func (*InternalTimeSeriesSample) ProtoMessage()    {}

func (m *InternalTimeSeriesSample) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *InternalTimeSeriesSample) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *InternalTimeSeriesSample) GetSum() float64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *InternalTimeSeriesSample) GetMax() float64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

func (m *InternalTimeSeriesSample) GetMin() float64 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

// RaftTruncatedState contains metadata about the truncated portion of the raft log.
// Raft requires access to the term of the last truncated log entry even after the
// rest of the entry has been discarded.
type RaftTruncatedState struct {
	// The highest index that has been removed from the log.
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index"`
	// The term corresponding to 'index'.
	Term uint64 `protobuf:"varint,2,opt,name=term" json:"term"`
}

func (m *RaftTruncatedState) Reset()         { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string { return proto1.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()    {}

func (m *RaftTruncatedState) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RaftTruncatedState) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

// RaftSnapshotData is the payload of a raftpb.Snapshot. It contains a raw copy of
// all of the range's data and metadata, including the raft log, response cache, etc.
type RaftSnapshotData struct {
	// The latest RangeDescriptor
	RangeDescriptor RangeDescriptor              `protobuf:"bytes,1,opt,name=range_descriptor" json:"range_descriptor"`
	KV              []*RaftSnapshotData_KeyValue `protobuf:"bytes,2,rep" json:"KV,omitempty"`
}

func (m *RaftSnapshotData) Reset()         { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string { return proto1.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()    {}

func (m *RaftSnapshotData) GetRangeDescriptor() RangeDescriptor {
	if m != nil {
		return m.RangeDescriptor
	}
	return RangeDescriptor{}
}

func (m *RaftSnapshotData) GetKV() []*RaftSnapshotData_KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type RaftSnapshotData_KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *RaftSnapshotData_KeyValue) Reset()         { *m = RaftSnapshotData_KeyValue{} }
func (m *RaftSnapshotData_KeyValue) String() string { return proto1.CompactTextString(m) }
func (*RaftSnapshotData_KeyValue) ProtoMessage()    {}

func (m *RaftSnapshotData_KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RaftSnapshotData_KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto1.RegisterEnum("cockroach.proto.InternalValueType", InternalValueType_name, InternalValueType_value)
}
func (m *RaftCommand) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftCommand) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.RangeID))
	data[i] = 0x10
	i++
	i = encodeVarintInternal(data, i, uint64(m.OriginNodeID))
	data[i] = 0x1a
	i++
	i = encodeVarintInternal(data, i, uint64(m.Cmd.Size()))
	n1, err := m.Cmd.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *InternalTimeSeriesData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalTimeSeriesData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.StartTimestampNanos))
	data[i] = 0x10
	i++
	i = encodeVarintInternal(data, i, uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, msg := range m.Samples {
			data[i] = 0x1a
			i++
			i = encodeVarintInternal(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *InternalTimeSeriesSample) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalTimeSeriesSample) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.Offset))
	data[i] = 0x30
	i++
	i = encodeVarintInternal(data, i, uint64(m.Count))
	data[i] = 0x39
	i++
	i = encodeFixed64Internal(data, i, uint64(math.Float64bits(m.Sum)))
	if m.Max != nil {
		data[i] = 0x41
		i++
		i = encodeFixed64Internal(data, i, uint64(math.Float64bits(*m.Max)))
	}
	if m.Min != nil {
		data[i] = 0x49
		i++
		i = encodeFixed64Internal(data, i, uint64(math.Float64bits(*m.Min)))
	}
	return i, nil
}

func (m *RaftTruncatedState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftTruncatedState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintInternal(data, i, uint64(m.Index))
	data[i] = 0x10
	i++
	i = encodeVarintInternal(data, i, uint64(m.Term))
	return i, nil
}

func (m *RaftSnapshotData) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftSnapshotData) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintInternal(data, i, uint64(m.RangeDescriptor.Size()))
	n2, err := m.RangeDescriptor.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.KV) > 0 {
		for _, msg := range m.KV {
			data[i] = 0x12
			i++
			i = encodeVarintInternal(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RaftSnapshotData_KeyValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftSnapshotData_KeyValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		data[i] = 0xa
		i++
		i = encodeVarintInternal(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if m.Value != nil {
		data[i] = 0x12
		i++
		i = encodeVarintInternal(data, i, uint64(len(m.Value)))
		i += copy(data[i:], m.Value)
	}
	return i, nil
}

func encodeFixed64Internal(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Internal(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInternal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RaftCommand) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.RangeID))
	n += 1 + sovInternal(uint64(m.OriginNodeID))
	l = m.Cmd.Size()
	n += 1 + l + sovInternal(uint64(l))
	return n
}

func (m *InternalTimeSeriesData) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.StartTimestampNanos))
	n += 1 + sovInternal(uint64(m.SampleDurationNanos))
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *InternalTimeSeriesSample) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.Offset))
	n += 1 + sovInternal(uint64(m.Count))
	n += 9
	if m.Max != nil {
		n += 9
	}
	if m.Min != nil {
		n += 9
	}
	return n
}

func (m *RaftTruncatedState) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovInternal(uint64(m.Index))
	n += 1 + sovInternal(uint64(m.Term))
	return n
}

func (m *RaftSnapshotData) Size() (n int) {
	var l int
	_ = l
	l = m.RangeDescriptor.Size()
	n += 1 + l + sovInternal(uint64(l))
	if len(m.KV) > 0 {
		for _, e := range m.KV {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *RaftSnapshotData_KeyValue) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovInternal(uint64(l))
	}
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovInternal(uint64(l))
	}
	return n
}

func sovInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternal(x uint64) (n int) {
	return sovInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftCommand) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeID |= (RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginNodeID", wireType)
			}
			m.OriginNodeID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.OriginNodeID |= (RaftNodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cmd.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *InternalTimeSeriesData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestampNanos", wireType)
			}
			m.StartTimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartTimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SampleDurationNanos", wireType)
			}
			m.SampleDurationNanos = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.SampleDurationNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Samples = append(m.Samples, &InternalTimeSeriesSample{})
			if err := m.Samples[len(m.Samples)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *InternalTimeSeriesSample) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Sum = float64(math.Float64frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Max = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Min = &v2
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RaftTruncatedState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RaftSnapshotData) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeDescriptor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RangeDescriptor.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KV", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KV = append(m.KV, &RaftSnapshotData_KeyValue{})
			if err := m.KV[len(m.KV)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RaftSnapshotData_KeyValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipInternal(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipInternal(data[start:])
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
	ErrInvalidLengthInternal = fmt.Errorf("proto: negative length found during unmarshaling")
)
