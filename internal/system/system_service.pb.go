// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/sgtsquiggs/appengine/internal/system/system_service.proto

package system

import proto "github.com/sgtsquiggs/protobuf/proto"
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

type SystemServiceError_ErrorCode int32

const (
	SystemServiceError_OK               SystemServiceError_ErrorCode = 0
	SystemServiceError_INTERNAL_ERROR   SystemServiceError_ErrorCode = 1
	SystemServiceError_BACKEND_REQUIRED SystemServiceError_ErrorCode = 2
	SystemServiceError_LIMIT_REACHED    SystemServiceError_ErrorCode = 3
)

var SystemServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "BACKEND_REQUIRED",
	3: "LIMIT_REACHED",
}
var SystemServiceError_ErrorCode_value = map[string]int32{
	"OK":               0,
	"INTERNAL_ERROR":   1,
	"BACKEND_REQUIRED": 2,
	"LIMIT_REACHED":    3,
}

func (x SystemServiceError_ErrorCode) Enum() *SystemServiceError_ErrorCode {
	p := new(SystemServiceError_ErrorCode)
	*p = x
	return p
}
func (x SystemServiceError_ErrorCode) String() string {
	return proto.EnumName(SystemServiceError_ErrorCode_name, int32(x))
}
func (x *SystemServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SystemServiceError_ErrorCode_value, data, "SystemServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = SystemServiceError_ErrorCode(value)
	return nil
}
func (SystemServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{0, 0}
}

type SystemServiceError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemServiceError) Reset()         { *m = SystemServiceError{} }
func (m *SystemServiceError) String() string { return proto.CompactTextString(m) }
func (*SystemServiceError) ProtoMessage()    {}
func (*SystemServiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{0}
}
func (m *SystemServiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemServiceError.Unmarshal(m, b)
}
func (m *SystemServiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemServiceError.Marshal(b, m, deterministic)
}
func (dst *SystemServiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemServiceError.Merge(dst, src)
}
func (m *SystemServiceError) XXX_Size() int {
	return xxx_messageInfo_SystemServiceError.Size(m)
}
func (m *SystemServiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemServiceError.DiscardUnknown(m)
}

var xxx_messageInfo_SystemServiceError proto.InternalMessageInfo

type SystemStat struct {
	// Instaneous value of this stat.
	Current *float64 `protobuf:"fixed64,1,opt,name=current" json:"current,omitempty"`
	// Average over time, if this stat has an instaneous value.
	Average1M  *float64 `protobuf:"fixed64,3,opt,name=average1m" json:"average1m,omitempty"`
	Average10M *float64 `protobuf:"fixed64,4,opt,name=average10m" json:"average10m,omitempty"`
	// Total value, if the stat accumulates over time.
	Total *float64 `protobuf:"fixed64,2,opt,name=total" json:"total,omitempty"`
	// Rate over time, if this stat accumulates.
	Rate1M               *float64 `protobuf:"fixed64,5,opt,name=rate1m" json:"rate1m,omitempty"`
	Rate10M              *float64 `protobuf:"fixed64,6,opt,name=rate10m" json:"rate10m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemStat) Reset()         { *m = SystemStat{} }
func (m *SystemStat) String() string { return proto.CompactTextString(m) }
func (*SystemStat) ProtoMessage()    {}
func (*SystemStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{1}
}
func (m *SystemStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStat.Unmarshal(m, b)
}
func (m *SystemStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStat.Marshal(b, m, deterministic)
}
func (dst *SystemStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStat.Merge(dst, src)
}
func (m *SystemStat) XXX_Size() int {
	return xxx_messageInfo_SystemStat.Size(m)
}
func (m *SystemStat) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStat.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStat proto.InternalMessageInfo

func (m *SystemStat) GetCurrent() float64 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *SystemStat) GetAverage1M() float64 {
	if m != nil && m.Average1M != nil {
		return *m.Average1M
	}
	return 0
}

func (m *SystemStat) GetAverage10M() float64 {
	if m != nil && m.Average10M != nil {
		return *m.Average10M
	}
	return 0
}

func (m *SystemStat) GetTotal() float64 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *SystemStat) GetRate1M() float64 {
	if m != nil && m.Rate1M != nil {
		return *m.Rate1M
	}
	return 0
}

func (m *SystemStat) GetRate10M() float64 {
	if m != nil && m.Rate10M != nil {
		return *m.Rate10M
	}
	return 0
}

type GetSystemStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSystemStatsRequest) Reset()         { *m = GetSystemStatsRequest{} }
func (m *GetSystemStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSystemStatsRequest) ProtoMessage()    {}
func (*GetSystemStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{2}
}
func (m *GetSystemStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSystemStatsRequest.Unmarshal(m, b)
}
func (m *GetSystemStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSystemStatsRequest.Marshal(b, m, deterministic)
}
func (dst *GetSystemStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSystemStatsRequest.Merge(dst, src)
}
func (m *GetSystemStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSystemStatsRequest.Size(m)
}
func (m *GetSystemStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSystemStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSystemStatsRequest proto.InternalMessageInfo

type GetSystemStatsResponse struct {
	// CPU used by this instance, in mcycles.
	Cpu *SystemStat `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	// Physical memory (RAM) used by this instance, in megabytes.
	Memory               *SystemStat `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetSystemStatsResponse) Reset()         { *m = GetSystemStatsResponse{} }
func (m *GetSystemStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSystemStatsResponse) ProtoMessage()    {}
func (*GetSystemStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{3}
}
func (m *GetSystemStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSystemStatsResponse.Unmarshal(m, b)
}
func (m *GetSystemStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSystemStatsResponse.Marshal(b, m, deterministic)
}
func (dst *GetSystemStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSystemStatsResponse.Merge(dst, src)
}
func (m *GetSystemStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSystemStatsResponse.Size(m)
}
func (m *GetSystemStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSystemStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSystemStatsResponse proto.InternalMessageInfo

func (m *GetSystemStatsResponse) GetCpu() *SystemStat {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *GetSystemStatsResponse) GetMemory() *SystemStat {
	if m != nil {
		return m.Memory
	}
	return nil
}

type StartBackgroundRequestRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackgroundRequestRequest) Reset()         { *m = StartBackgroundRequestRequest{} }
func (m *StartBackgroundRequestRequest) String() string { return proto.CompactTextString(m) }
func (*StartBackgroundRequestRequest) ProtoMessage()    {}
func (*StartBackgroundRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{4}
}
func (m *StartBackgroundRequestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackgroundRequestRequest.Unmarshal(m, b)
}
func (m *StartBackgroundRequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackgroundRequestRequest.Marshal(b, m, deterministic)
}
func (dst *StartBackgroundRequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackgroundRequestRequest.Merge(dst, src)
}
func (m *StartBackgroundRequestRequest) XXX_Size() int {
	return xxx_messageInfo_StartBackgroundRequestRequest.Size(m)
}
func (m *StartBackgroundRequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackgroundRequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackgroundRequestRequest proto.InternalMessageInfo

type StartBackgroundRequestResponse struct {
	// Every /_ah/background request will have an X-AppEngine-BackgroundRequest
	// header, whose value will be equal to this parameter, the request_id.
	RequestId            *string  `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackgroundRequestResponse) Reset()         { *m = StartBackgroundRequestResponse{} }
func (m *StartBackgroundRequestResponse) String() string { return proto.CompactTextString(m) }
func (*StartBackgroundRequestResponse) ProtoMessage()    {}
func (*StartBackgroundRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_system_service_ccf41ec210fc59eb, []int{5}
}
func (m *StartBackgroundRequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackgroundRequestResponse.Unmarshal(m, b)
}
func (m *StartBackgroundRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackgroundRequestResponse.Marshal(b, m, deterministic)
}
func (dst *StartBackgroundRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackgroundRequestResponse.Merge(dst, src)
}
func (m *StartBackgroundRequestResponse) XXX_Size() int {
	return xxx_messageInfo_StartBackgroundRequestResponse.Size(m)
}
func (m *StartBackgroundRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackgroundRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackgroundRequestResponse proto.InternalMessageInfo

func (m *StartBackgroundRequestResponse) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemServiceError)(nil), "appengine.SystemServiceError")
	proto.RegisterType((*SystemStat)(nil), "appengine.SystemStat")
	proto.RegisterType((*GetSystemStatsRequest)(nil), "appengine.GetSystemStatsRequest")
	proto.RegisterType((*GetSystemStatsResponse)(nil), "appengine.GetSystemStatsResponse")
	proto.RegisterType((*StartBackgroundRequestRequest)(nil), "appengine.StartBackgroundRequestRequest")
	proto.RegisterType((*StartBackgroundRequestResponse)(nil), "appengine.StartBackgroundRequestResponse")
}

func init() {
	proto.RegisterFile("github.com/sgtsquiggs/appengine/internal/system/system_service.proto", fileDescriptor_system_service_ccf41ec210fc59eb)
}

var fileDescriptor_system_service_ccf41ec210fc59eb = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8f, 0x93, 0x40,
	0x18, 0xc6, 0xa5, 0x75, 0x51, 0x5e, 0xa3, 0xc1, 0xc9, 0xee, 0xca, 0xc1, 0x5d, 0x0d, 0x17, 0xbd,
	0x48, 0x57, 0xbf, 0x80, 0xf6, 0xcf, 0x44, 0x49, 0x6b, 0xab, 0xd3, 0x7a, 0xf1, 0x42, 0x26, 0xf0,
	0x3a, 0x21, 0xc2, 0x0c, 0x0e, 0x43, 0x93, 0x7e, 0x27, 0x3f, 0xa4, 0xe9, 0x30, 0x6d, 0xcd, 0x26,
	0x3d, 0x31, 0xcf, 0xf3, 0xfc, 0x02, 0x3f, 0x08, 0xf0, 0x49, 0x28, 0x25, 0x2a, 0x4c, 0x84, 0xaa,
	0xb8, 0x14, 0x89, 0xd2, 0x62, 0xc4, 0x9b, 0x06, 0xa5, 0x28, 0x25, 0x8e, 0x4a, 0x69, 0x50, 0x4b,
	0x5e, 0x8d, 0xda, 0x5d, 0x6b, 0xb0, 0x76, 0x97, 0xac, 0x45, 0xbd, 0x2d, 0x73, 0x4c, 0x1a, 0xad,
	0x8c, 0x22, 0xc1, 0x91, 0x8f, 0x7f, 0x01, 0x59, 0x5b, 0x64, 0xdd, 0x13, 0x54, 0x6b, 0xa5, 0xe3,
	0x6f, 0x10, 0xd8, 0xc3, 0x54, 0x15, 0x48, 0x7c, 0x18, 0xac, 0xe6, 0xe1, 0x03, 0x42, 0xe0, 0x59,
	0xba, 0xdc, 0x50, 0xb6, 0x1c, 0x2f, 0x32, 0xca, 0xd8, 0x8a, 0x85, 0x1e, 0xb9, 0x84, 0x70, 0x32,
	0x9e, 0xce, 0xe9, 0x72, 0x96, 0x31, 0xfa, 0xfd, 0x47, 0xca, 0xe8, 0x2c, 0x1c, 0x90, 0xe7, 0xf0,
	0x74, 0x91, 0x7e, 0x4d, 0x37, 0x19, 0xa3, 0xe3, 0xe9, 0x17, 0x3a, 0x0b, 0x87, 0xf1, 0x5f, 0x0f,
	0xc0, 0x3d, 0xc8, 0x70, 0x43, 0x22, 0x78, 0x94, 0x77, 0x5a, 0xa3, 0x34, 0x91, 0xf7, 0xda, 0x7b,
	0xeb, 0xb1, 0x43, 0x24, 0x2f, 0x21, 0xe0, 0x5b, 0xd4, 0x5c, 0xe0, 0xfb, 0x3a, 0x1a, 0xda, 0xed,
	0x54, 0x90, 0x5b, 0x80, 0x43, 0xb8, 0xab, 0xa3, 0x87, 0x76, 0xfe, 0xaf, 0x21, 0x97, 0x70, 0x61,
	0x94, 0xe1, 0x55, 0x34, 0xb0, 0x53, 0x1f, 0xc8, 0x35, 0xf8, 0x9a, 0x9b, 0xfd, 0x0d, 0x2f, 0x6c,
	0xed, 0xd2, 0xde, 0xc2, 0x9e, 0xee, 0xea, 0xc8, 0xef, 0x2d, 0x5c, 0x8c, 0x5f, 0xc0, 0xd5, 0x67,
	0x34, 0x27, 0xe1, 0x96, 0xe1, 0x9f, 0x0e, 0x5b, 0x13, 0x37, 0x70, 0x7d, 0x7f, 0x68, 0x1b, 0x25,
	0x5b, 0x24, 0x6f, 0x60, 0x98, 0x37, 0x9d, 0x7d, 0x9d, 0x27, 0x1f, 0xae, 0x92, 0xe3, 0x27, 0x4e,
	0x4e, 0x30, 0xdb, 0x13, 0xe4, 0x1d, 0xf8, 0x35, 0xd6, 0x4a, 0xef, 0xac, 0xe4, 0x59, 0xd6, 0x41,
	0xf1, 0x2b, 0xb8, 0x59, 0x1b, 0xae, 0xcd, 0x84, 0xe7, 0xbf, 0x85, 0x56, 0x9d, 0x2c, 0x9c, 0xcb,
	0x41, 0xe9, 0x23, 0xdc, 0x9e, 0x03, 0x9c, 0xda, 0x0d, 0x80, 0xee, 0xab, 0xac, 0x2c, 0xac, 0x61,
	0xc0, 0x02, 0xd7, 0xa4, 0xc5, 0xe4, 0xf1, 0x4f, 0xbf, 0xff, 0x4d, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x56, 0x5d, 0x5e, 0xc3, 0x5b, 0x02, 0x00, 0x00,
}
