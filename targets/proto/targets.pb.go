// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/proto/targets.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/targets/proto/targets.proto

It has these top-level messages:
	TargetsDef
	DummyTargets
	GlobalTargetsOptions
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_targets_lameduck "github.com/google/cloudprober/targets/lameduck/proto"
import cloudprober_targets_gce "github.com/google/cloudprober/targets/gce/proto"
import cloudprober_targets_rtc "github.com/google/cloudprober/targets/rtc/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type TargetsDef struct {
	// Types that are valid to be assigned to Type:
	//	*TargetsDef_HostNames
	//	*TargetsDef_GceTargets
	//	*TargetsDef_RtcTargets
	//	*TargetsDef_DummyTargets
	Type isTargetsDef_Type `protobuf_oneof:"type"`
	// Regex to apply on the targets.
	Regex *string `protobuf:"bytes,4,opt,name=regex" json:"regex,omitempty"`
	// Exclude lameducks. Lameduck targets can be set through RTC (realtime
	// configurator) service. This functionality works only if lame_duck_options
	// are specified.
	ExcludeLameducks *bool `protobuf:"varint,5,opt,name=exclude_lameducks,json=excludeLameducks,def=1" json:"exclude_lameducks,omitempty"`
	// How often targets should be evaluated. Any number less than or equal to 0
	// will result in no target caching (targets will be reevaluated on demand).
	// Note that individual target types may have their own caches implemented
	// (specifically GCE instances/forwarding rules). This does not impact those
	// caches.
	ReEvalSec                     *int32 `protobuf:"varint,6,opt,name=re_eval_sec,json=reEvalSec,def=0" json:"re_eval_sec,omitempty"`
	proto1.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized              []byte `json:"-"`
}

func (m *TargetsDef) Reset()                    { *m = TargetsDef{} }
func (m *TargetsDef) String() string            { return proto1.CompactTextString(m) }
func (*TargetsDef) ProtoMessage()               {}
func (*TargetsDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_TargetsDef = []proto1.ExtensionRange{
	{200, 536870911},
}

func (*TargetsDef) ExtensionRangeArray() []proto1.ExtensionRange {
	return extRange_TargetsDef
}

const Default_TargetsDef_ExcludeLameducks bool = true
const Default_TargetsDef_ReEvalSec int32 = 0

type isTargetsDef_Type interface {
	isTargetsDef_Type()
}

type TargetsDef_HostNames struct {
	HostNames string `protobuf:"bytes,1,opt,name=host_names,json=hostNames,oneof"`
}
type TargetsDef_GceTargets struct {
	GceTargets *cloudprober_targets_gce.TargetsConf `protobuf:"bytes,2,opt,name=gce_targets,json=gceTargets,oneof"`
}
type TargetsDef_RtcTargets struct {
	RtcTargets *cloudprober_targets_rtc.TargetsConf `protobuf:"bytes,3,opt,name=rtc_targets,json=rtcTargets,oneof"`
}
type TargetsDef_DummyTargets struct {
	DummyTargets *DummyTargets `protobuf:"bytes,7,opt,name=dummy_targets,json=dummyTargets,oneof"`
}

func (*TargetsDef_HostNames) isTargetsDef_Type()    {}
func (*TargetsDef_GceTargets) isTargetsDef_Type()   {}
func (*TargetsDef_RtcTargets) isTargetsDef_Type()   {}
func (*TargetsDef_DummyTargets) isTargetsDef_Type() {}

func (m *TargetsDef) GetType() isTargetsDef_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *TargetsDef) GetHostNames() string {
	if x, ok := m.GetType().(*TargetsDef_HostNames); ok {
		return x.HostNames
	}
	return ""
}

func (m *TargetsDef) GetGceTargets() *cloudprober_targets_gce.TargetsConf {
	if x, ok := m.GetType().(*TargetsDef_GceTargets); ok {
		return x.GceTargets
	}
	return nil
}

func (m *TargetsDef) GetRtcTargets() *cloudprober_targets_rtc.TargetsConf {
	if x, ok := m.GetType().(*TargetsDef_RtcTargets); ok {
		return x.RtcTargets
	}
	return nil
}

func (m *TargetsDef) GetDummyTargets() *DummyTargets {
	if x, ok := m.GetType().(*TargetsDef_DummyTargets); ok {
		return x.DummyTargets
	}
	return nil
}

func (m *TargetsDef) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *TargetsDef) GetExcludeLameducks() bool {
	if m != nil && m.ExcludeLameducks != nil {
		return *m.ExcludeLameducks
	}
	return Default_TargetsDef_ExcludeLameducks
}

func (m *TargetsDef) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_TargetsDef_ReEvalSec
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TargetsDef) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _TargetsDef_OneofMarshaler, _TargetsDef_OneofUnmarshaler, _TargetsDef_OneofSizer, []interface{}{
		(*TargetsDef_HostNames)(nil),
		(*TargetsDef_GceTargets)(nil),
		(*TargetsDef_RtcTargets)(nil),
		(*TargetsDef_DummyTargets)(nil),
	}
}

func _TargetsDef_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*TargetsDef)
	// type
	switch x := m.Type.(type) {
	case *TargetsDef_HostNames:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.HostNames)
	case *TargetsDef_GceTargets:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.GceTargets); err != nil {
			return err
		}
	case *TargetsDef_RtcTargets:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.RtcTargets); err != nil {
			return err
		}
	case *TargetsDef_DummyTargets:
		b.EncodeVarint(7<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.DummyTargets); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TargetsDef.Type has unexpected type %T", x)
	}
	return nil
}

func _TargetsDef_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*TargetsDef)
	switch tag {
	case 1: // type.host_names
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &TargetsDef_HostNames{x}
		return true, err
	case 2: // type.gce_targets
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(cloudprober_targets_gce.TargetsConf)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_GceTargets{msg}
		return true, err
	case 3: // type.rtc_targets
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(cloudprober_targets_rtc.TargetsConf)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_RtcTargets{msg}
		return true, err
	case 7: // type.dummy_targets
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(DummyTargets)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_DummyTargets{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TargetsDef_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*TargetsDef)
	// type
	switch x := m.Type.(type) {
	case *TargetsDef_HostNames:
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(len(x.HostNames)))
		n += len(x.HostNames)
	case *TargetsDef_GceTargets:
		s := proto1.Size(x.GceTargets)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *TargetsDef_RtcTargets:
		s := proto1.Size(x.RtcTargets)
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *TargetsDef_DummyTargets:
		s := proto1.Size(x.DummyTargets)
		n += proto1.SizeVarint(7<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// DummyTargets represent empty targets, which are useful for external
// probes that do not have any "proper" targets.  Such as ilbprober.
type DummyTargets struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DummyTargets) Reset()                    { *m = DummyTargets{} }
func (m *DummyTargets) String() string            { return proto1.CompactTextString(m) }
func (*DummyTargets) ProtoMessage()               {}
func (*DummyTargets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Global targets options. These options are independent of the per-probe
// targets which are defined by the "Targets" type above.
//
// Currently these options are used only for GCE targets to control things like
// how often to re-evaluate the targets and whether to check for lame ducks or
// not.
type GlobalTargetsOptions struct {
	// GCE targets options.
	GlobalGceTargetsOptions *cloudprober_targets_gce.GlobalOptions `protobuf:"bytes,1,opt,name=global_gce_targets_options,json=globalGceTargetsOptions" json:"global_gce_targets_options,omitempty"`
	// Lame duck options. If provided, targets module checks for the lame duck
	// targets and removes them from the targets list.
	LameDuckOptions  *cloudprober_targets_lameduck.Options `protobuf:"bytes,2,opt,name=lame_duck_options,json=lameDuckOptions" json:"lame_duck_options,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *GlobalTargetsOptions) Reset()                    { *m = GlobalTargetsOptions{} }
func (m *GlobalTargetsOptions) String() string            { return proto1.CompactTextString(m) }
func (*GlobalTargetsOptions) ProtoMessage()               {}
func (*GlobalTargetsOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GlobalTargetsOptions) GetGlobalGceTargetsOptions() *cloudprober_targets_gce.GlobalOptions {
	if m != nil {
		return m.GlobalGceTargetsOptions
	}
	return nil
}

func (m *GlobalTargetsOptions) GetLameDuckOptions() *cloudprober_targets_lameduck.Options {
	if m != nil {
		return m.LameDuckOptions
	}
	return nil
}

func init() {
	proto1.RegisterType((*TargetsDef)(nil), "cloudprober.targets.TargetsDef")
	proto1.RegisterType((*DummyTargets)(nil), "cloudprober.targets.DummyTargets")
	proto1.RegisterType((*GlobalTargetsOptions)(nil), "cloudprober.targets.GlobalTargetsOptions")
}

func init() {
	proto1.RegisterFile("github.com/google/cloudprober/targets/proto/targets.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd1, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x07, 0xf0, 0x7a, 0x6b, 0xc7, 0xfa, 0x3a, 0x60, 0x0b, 0x93, 0x88, 0x7a, 0xa1, 0xab, 0x00,
	0x55, 0x1c, 0x52, 0xd8, 0x8d, 0x89, 0x0b, 0x50, 0xd4, 0x1e, 0x10, 0x88, 0xc0, 0xdd, 0x4a, 0x5f,
	0x5e, 0xbd, 0x6a, 0x4e, 0x1d, 0x39, 0xce, 0xb4, 0xdd, 0x72, 0xe4, 0xa3, 0x71, 0xe2, 0x33, 0x21,
	0x27, 0x4e, 0x14, 0xa1, 0x80, 0x26, 0xed, 0xf8, 0xfe, 0xcf, 0xfe, 0x25, 0xcf, 0x0f, 0xde, 0x8a,
	0xad, 0xb9, 0xcc, 0xd7, 0x01, 0xaa, 0x64, 0x2e, 0x94, 0x12, 0x92, 0xe6, 0x28, 0x55, 0x1e, 0xa7,
	0x5a, 0xad, 0x49, 0xcf, 0x4d, 0xa4, 0x05, 0x99, 0x6c, 0x9e, 0x6a, 0x65, 0x54, 0x5d, 0x05, 0x65,
	0xe5, 0x3d, 0x69, 0x1d, 0x0c, 0x5c, 0x6b, 0xfc, 0xfe, 0x6e, 0x9e, 0x8c, 0x12, 0x8a, 0x73, 0xbc,
	0x72, 0x30, 0xaa, 0xdd, 0x66, 0x2b, 0x2a, 0x77, 0xfc, 0xee, 0x6e, 0x84, 0x40, 0xba, 0xc7, 0x6d,
	0x6d, 0xb0, 0xe3, 0xf6, 0xf4, 0xe7, 0x3e, 0xc0, 0x8f, 0xea, 0xc8, 0x82, 0x36, 0xde, 0x33, 0x80,
	0x4b, 0x95, 0x19, 0xbe, 0x8b, 0x12, 0xca, 0x7c, 0x36, 0x61, 0xb3, 0xe1, 0xaa, 0x17, 0x0e, 0x6d,
	0xf6, 0xc5, 0x46, 0xde, 0x12, 0x46, 0x02, 0x89, 0x3b, 0xd5, 0xdf, 0x9b, 0xb0, 0xd9, 0xe8, 0xfc,
	0x79, 0xd0, 0xf1, 0x32, 0x81, 0x40, 0x0a, 0x1c, 0xfd, 0x51, 0xed, 0x36, 0xab, 0x5e, 0x08, 0x02,
	0xc9, 0x25, 0x16, 0xd2, 0x06, 0x1b, 0x68, 0xff, 0x3f, 0x90, 0x36, 0xf8, 0x37, 0xa4, 0x0d, 0xd6,
	0xd0, 0x0a, 0x1e, 0xc6, 0x79, 0x92, 0xdc, 0x36, 0xd4, 0x83, 0x92, 0x3a, 0xeb, 0xa4, 0x16, 0xf6,
	0xa4, 0xbb, 0xb9, 0xea, 0x85, 0x47, 0x71, 0xab, 0xf6, 0x4e, 0x61, 0xa0, 0x49, 0xd0, 0x8d, 0xdf,
	0xb7, 0x73, 0x87, 0x55, 0xe1, 0xbd, 0x81, 0x13, 0xba, 0x41, 0x99, 0xc7, 0xc4, 0xeb, 0x25, 0x66,
	0xfe, 0x60, 0xc2, 0x66, 0x87, 0x17, 0x7d, 0xa3, 0x73, 0x0a, 0x8f, 0x5d, 0xfb, 0x73, 0xdd, 0xf5,
	0xce, 0x60, 0xa4, 0x89, 0xd3, 0x75, 0x24, 0x79, 0x46, 0xe8, 0x1f, 0x4c, 0xd8, 0x6c, 0x70, 0xc1,
	0x5e, 0x87, 0x43, 0x4d, 0x9f, 0xae, 0x23, 0xf9, 0x9d, 0xf0, 0xd5, 0xf0, 0xf0, 0x17, 0x3b, 0x2e,
	0x8a, 0xa2, 0xd8, 0xfb, 0x70, 0x00, 0x7d, 0x73, 0x9b, 0xd2, 0xf4, 0x11, 0x1c, 0xb5, 0x7f, 0x6f,
	0xfa, 0x9b, 0xc1, 0xe9, 0x52, 0xaa, 0x75, 0x24, 0x5d, 0xf2, 0x35, 0x35, 0x5b, 0xb5, 0xcb, 0x3c,
	0x84, 0xb1, 0x28, 0x73, 0xde, 0x5a, 0x05, 0x57, 0x55, 0xb7, 0x5c, 0xda, 0xe8, 0xfc, 0xe5, 0x3f,
	0x57, 0x52, 0x91, 0xce, 0x0a, 0x9f, 0x56, 0xd2, 0xb2, 0x59, 0x4c, 0xfd, 0x91, 0x6f, 0x70, 0x62,
	0xc7, 0xe5, 0x76, 0xa2, 0xc6, 0xae, 0xd6, 0xfd, 0xa2, 0xd3, 0xae, 0x1f, 0x27, 0xa8, 0xe9, 0xc7,
	0x36, 0x59, 0xe4, 0x78, 0xe5, 0x82, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x5d, 0xf1, 0x0b,
	0x7b, 0x03, 0x00, 0x00,
}
