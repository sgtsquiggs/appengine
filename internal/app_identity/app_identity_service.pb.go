// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/sgtsquiggs/appengine/internal/app_identity/app_identity_service.proto

package app_identity

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

type AppIdentityServiceError_ErrorCode int32

const (
	AppIdentityServiceError_SUCCESS           AppIdentityServiceError_ErrorCode = 0
	AppIdentityServiceError_UNKNOWN_SCOPE     AppIdentityServiceError_ErrorCode = 9
	AppIdentityServiceError_BLOB_TOO_LARGE    AppIdentityServiceError_ErrorCode = 1000
	AppIdentityServiceError_DEADLINE_EXCEEDED AppIdentityServiceError_ErrorCode = 1001
	AppIdentityServiceError_NOT_A_VALID_APP   AppIdentityServiceError_ErrorCode = 1002
	AppIdentityServiceError_UNKNOWN_ERROR     AppIdentityServiceError_ErrorCode = 1003
	AppIdentityServiceError_NOT_ALLOWED       AppIdentityServiceError_ErrorCode = 1005
	AppIdentityServiceError_NOT_IMPLEMENTED   AppIdentityServiceError_ErrorCode = 1006
)

var AppIdentityServiceError_ErrorCode_name = map[int32]string{
	0:    "SUCCESS",
	9:    "UNKNOWN_SCOPE",
	1000: "BLOB_TOO_LARGE",
	1001: "DEADLINE_EXCEEDED",
	1002: "NOT_A_VALID_APP",
	1003: "UNKNOWN_ERROR",
	1005: "NOT_ALLOWED",
	1006: "NOT_IMPLEMENTED",
}
var AppIdentityServiceError_ErrorCode_value = map[string]int32{
	"SUCCESS":           0,
	"UNKNOWN_SCOPE":     9,
	"BLOB_TOO_LARGE":    1000,
	"DEADLINE_EXCEEDED": 1001,
	"NOT_A_VALID_APP":   1002,
	"UNKNOWN_ERROR":     1003,
	"NOT_ALLOWED":       1005,
	"NOT_IMPLEMENTED":   1006,
}

func (x AppIdentityServiceError_ErrorCode) Enum() *AppIdentityServiceError_ErrorCode {
	p := new(AppIdentityServiceError_ErrorCode)
	*p = x
	return p
}
func (x AppIdentityServiceError_ErrorCode) String() string {
	return proto.EnumName(AppIdentityServiceError_ErrorCode_name, int32(x))
}
func (x *AppIdentityServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AppIdentityServiceError_ErrorCode_value, data, "AppIdentityServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = AppIdentityServiceError_ErrorCode(value)
	return nil
}
func (AppIdentityServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{0, 0}
}

type AppIdentityServiceError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppIdentityServiceError) Reset()         { *m = AppIdentityServiceError{} }
func (m *AppIdentityServiceError) String() string { return proto.CompactTextString(m) }
func (*AppIdentityServiceError) ProtoMessage()    {}
func (*AppIdentityServiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{0}
}
func (m *AppIdentityServiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppIdentityServiceError.Unmarshal(m, b)
}
func (m *AppIdentityServiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppIdentityServiceError.Marshal(b, m, deterministic)
}
func (dst *AppIdentityServiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppIdentityServiceError.Merge(dst, src)
}
func (m *AppIdentityServiceError) XXX_Size() int {
	return xxx_messageInfo_AppIdentityServiceError.Size(m)
}
func (m *AppIdentityServiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_AppIdentityServiceError.DiscardUnknown(m)
}

var xxx_messageInfo_AppIdentityServiceError proto.InternalMessageInfo

type SignForAppRequest struct {
	BytesToSign          []byte   `protobuf:"bytes,1,opt,name=bytes_to_sign,json=bytesToSign" json:"bytes_to_sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignForAppRequest) Reset()         { *m = SignForAppRequest{} }
func (m *SignForAppRequest) String() string { return proto.CompactTextString(m) }
func (*SignForAppRequest) ProtoMessage()    {}
func (*SignForAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{1}
}
func (m *SignForAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignForAppRequest.Unmarshal(m, b)
}
func (m *SignForAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignForAppRequest.Marshal(b, m, deterministic)
}
func (dst *SignForAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignForAppRequest.Merge(dst, src)
}
func (m *SignForAppRequest) XXX_Size() int {
	return xxx_messageInfo_SignForAppRequest.Size(m)
}
func (m *SignForAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignForAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignForAppRequest proto.InternalMessageInfo

func (m *SignForAppRequest) GetBytesToSign() []byte {
	if m != nil {
		return m.BytesToSign
	}
	return nil
}

type SignForAppResponse struct {
	KeyName              *string  `protobuf:"bytes,1,opt,name=key_name,json=keyName" json:"key_name,omitempty"`
	SignatureBytes       []byte   `protobuf:"bytes,2,opt,name=signature_bytes,json=signatureBytes" json:"signature_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignForAppResponse) Reset()         { *m = SignForAppResponse{} }
func (m *SignForAppResponse) String() string { return proto.CompactTextString(m) }
func (*SignForAppResponse) ProtoMessage()    {}
func (*SignForAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{2}
}
func (m *SignForAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignForAppResponse.Unmarshal(m, b)
}
func (m *SignForAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignForAppResponse.Marshal(b, m, deterministic)
}
func (dst *SignForAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignForAppResponse.Merge(dst, src)
}
func (m *SignForAppResponse) XXX_Size() int {
	return xxx_messageInfo_SignForAppResponse.Size(m)
}
func (m *SignForAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignForAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignForAppResponse proto.InternalMessageInfo

func (m *SignForAppResponse) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

func (m *SignForAppResponse) GetSignatureBytes() []byte {
	if m != nil {
		return m.SignatureBytes
	}
	return nil
}

type GetPublicCertificateForAppRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicCertificateForAppRequest) Reset()         { *m = GetPublicCertificateForAppRequest{} }
func (m *GetPublicCertificateForAppRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicCertificateForAppRequest) ProtoMessage()    {}
func (*GetPublicCertificateForAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{3}
}
func (m *GetPublicCertificateForAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicCertificateForAppRequest.Unmarshal(m, b)
}
func (m *GetPublicCertificateForAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicCertificateForAppRequest.Marshal(b, m, deterministic)
}
func (dst *GetPublicCertificateForAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicCertificateForAppRequest.Merge(dst, src)
}
func (m *GetPublicCertificateForAppRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicCertificateForAppRequest.Size(m)
}
func (m *GetPublicCertificateForAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicCertificateForAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicCertificateForAppRequest proto.InternalMessageInfo

type PublicCertificate struct {
	KeyName              *string  `protobuf:"bytes,1,opt,name=key_name,json=keyName" json:"key_name,omitempty"`
	X509CertificatePem   *string  `protobuf:"bytes,2,opt,name=x509_certificate_pem,json=x509CertificatePem" json:"x509_certificate_pem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicCertificate) Reset()         { *m = PublicCertificate{} }
func (m *PublicCertificate) String() string { return proto.CompactTextString(m) }
func (*PublicCertificate) ProtoMessage()    {}
func (*PublicCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{4}
}
func (m *PublicCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicCertificate.Unmarshal(m, b)
}
func (m *PublicCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicCertificate.Marshal(b, m, deterministic)
}
func (dst *PublicCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicCertificate.Merge(dst, src)
}
func (m *PublicCertificate) XXX_Size() int {
	return xxx_messageInfo_PublicCertificate.Size(m)
}
func (m *PublicCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_PublicCertificate proto.InternalMessageInfo

func (m *PublicCertificate) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

func (m *PublicCertificate) GetX509CertificatePem() string {
	if m != nil && m.X509CertificatePem != nil {
		return *m.X509CertificatePem
	}
	return ""
}

type GetPublicCertificateForAppResponse struct {
	PublicCertificateList      []*PublicCertificate `protobuf:"bytes,1,rep,name=public_certificate_list,json=publicCertificateList" json:"public_certificate_list,omitempty"`
	MaxClientCacheTimeInSecond *int64               `protobuf:"varint,2,opt,name=max_client_cache_time_in_second,json=maxClientCacheTimeInSecond" json:"max_client_cache_time_in_second,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *GetPublicCertificateForAppResponse) Reset()         { *m = GetPublicCertificateForAppResponse{} }
func (m *GetPublicCertificateForAppResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicCertificateForAppResponse) ProtoMessage()    {}
func (*GetPublicCertificateForAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{5}
}
func (m *GetPublicCertificateForAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicCertificateForAppResponse.Unmarshal(m, b)
}
func (m *GetPublicCertificateForAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicCertificateForAppResponse.Marshal(b, m, deterministic)
}
func (dst *GetPublicCertificateForAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicCertificateForAppResponse.Merge(dst, src)
}
func (m *GetPublicCertificateForAppResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicCertificateForAppResponse.Size(m)
}
func (m *GetPublicCertificateForAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicCertificateForAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicCertificateForAppResponse proto.InternalMessageInfo

func (m *GetPublicCertificateForAppResponse) GetPublicCertificateList() []*PublicCertificate {
	if m != nil {
		return m.PublicCertificateList
	}
	return nil
}

func (m *GetPublicCertificateForAppResponse) GetMaxClientCacheTimeInSecond() int64 {
	if m != nil && m.MaxClientCacheTimeInSecond != nil {
		return *m.MaxClientCacheTimeInSecond
	}
	return 0
}

type GetServiceAccountNameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceAccountNameRequest) Reset()         { *m = GetServiceAccountNameRequest{} }
func (m *GetServiceAccountNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceAccountNameRequest) ProtoMessage()    {}
func (*GetServiceAccountNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{6}
}
func (m *GetServiceAccountNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceAccountNameRequest.Unmarshal(m, b)
}
func (m *GetServiceAccountNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceAccountNameRequest.Marshal(b, m, deterministic)
}
func (dst *GetServiceAccountNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceAccountNameRequest.Merge(dst, src)
}
func (m *GetServiceAccountNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceAccountNameRequest.Size(m)
}
func (m *GetServiceAccountNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceAccountNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceAccountNameRequest proto.InternalMessageInfo

type GetServiceAccountNameResponse struct {
	ServiceAccountName   *string  `protobuf:"bytes,1,opt,name=service_account_name,json=serviceAccountName" json:"service_account_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceAccountNameResponse) Reset()         { *m = GetServiceAccountNameResponse{} }
func (m *GetServiceAccountNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceAccountNameResponse) ProtoMessage()    {}
func (*GetServiceAccountNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{7}
}
func (m *GetServiceAccountNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceAccountNameResponse.Unmarshal(m, b)
}
func (m *GetServiceAccountNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceAccountNameResponse.Marshal(b, m, deterministic)
}
func (dst *GetServiceAccountNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceAccountNameResponse.Merge(dst, src)
}
func (m *GetServiceAccountNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceAccountNameResponse.Size(m)
}
func (m *GetServiceAccountNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceAccountNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceAccountNameResponse proto.InternalMessageInfo

func (m *GetServiceAccountNameResponse) GetServiceAccountName() string {
	if m != nil && m.ServiceAccountName != nil {
		return *m.ServiceAccountName
	}
	return ""
}

type GetAccessTokenRequest struct {
	Scope                []string `protobuf:"bytes,1,rep,name=scope" json:"scope,omitempty"`
	ServiceAccountId     *int64   `protobuf:"varint,2,opt,name=service_account_id,json=serviceAccountId" json:"service_account_id,omitempty"`
	ServiceAccountName   *string  `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName" json:"service_account_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccessTokenRequest) Reset()         { *m = GetAccessTokenRequest{} }
func (m *GetAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccessTokenRequest) ProtoMessage()    {}
func (*GetAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{8}
}
func (m *GetAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccessTokenRequest.Unmarshal(m, b)
}
func (m *GetAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccessTokenRequest.Marshal(b, m, deterministic)
}
func (dst *GetAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessTokenRequest.Merge(dst, src)
}
func (m *GetAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccessTokenRequest.Size(m)
}
func (m *GetAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessTokenRequest proto.InternalMessageInfo

func (m *GetAccessTokenRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *GetAccessTokenRequest) GetServiceAccountId() int64 {
	if m != nil && m.ServiceAccountId != nil {
		return *m.ServiceAccountId
	}
	return 0
}

func (m *GetAccessTokenRequest) GetServiceAccountName() string {
	if m != nil && m.ServiceAccountName != nil {
		return *m.ServiceAccountName
	}
	return ""
}

type GetAccessTokenResponse struct {
	AccessToken          *string  `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	ExpirationTime       *int64   `protobuf:"varint,2,opt,name=expiration_time,json=expirationTime" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccessTokenResponse) Reset()         { *m = GetAccessTokenResponse{} }
func (m *GetAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccessTokenResponse) ProtoMessage()    {}
func (*GetAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{9}
}
func (m *GetAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccessTokenResponse.Unmarshal(m, b)
}
func (m *GetAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccessTokenResponse.Marshal(b, m, deterministic)
}
func (dst *GetAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessTokenResponse.Merge(dst, src)
}
func (m *GetAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccessTokenResponse.Size(m)
}
func (m *GetAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessTokenResponse proto.InternalMessageInfo

func (m *GetAccessTokenResponse) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *GetAccessTokenResponse) GetExpirationTime() int64 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return 0
}

type GetDefaultGcsBucketNameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefaultGcsBucketNameRequest) Reset()         { *m = GetDefaultGcsBucketNameRequest{} }
func (m *GetDefaultGcsBucketNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefaultGcsBucketNameRequest) ProtoMessage()    {}
func (*GetDefaultGcsBucketNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{10}
}
func (m *GetDefaultGcsBucketNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefaultGcsBucketNameRequest.Unmarshal(m, b)
}
func (m *GetDefaultGcsBucketNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefaultGcsBucketNameRequest.Marshal(b, m, deterministic)
}
func (dst *GetDefaultGcsBucketNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefaultGcsBucketNameRequest.Merge(dst, src)
}
func (m *GetDefaultGcsBucketNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetDefaultGcsBucketNameRequest.Size(m)
}
func (m *GetDefaultGcsBucketNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefaultGcsBucketNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefaultGcsBucketNameRequest proto.InternalMessageInfo

type GetDefaultGcsBucketNameResponse struct {
	DefaultGcsBucketName *string  `protobuf:"bytes,1,opt,name=default_gcs_bucket_name,json=defaultGcsBucketName" json:"default_gcs_bucket_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefaultGcsBucketNameResponse) Reset()         { *m = GetDefaultGcsBucketNameResponse{} }
func (m *GetDefaultGcsBucketNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetDefaultGcsBucketNameResponse) ProtoMessage()    {}
func (*GetDefaultGcsBucketNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_identity_service_08a6e3f74b04cfa4, []int{11}
}
func (m *GetDefaultGcsBucketNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefaultGcsBucketNameResponse.Unmarshal(m, b)
}
func (m *GetDefaultGcsBucketNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefaultGcsBucketNameResponse.Marshal(b, m, deterministic)
}
func (dst *GetDefaultGcsBucketNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefaultGcsBucketNameResponse.Merge(dst, src)
}
func (m *GetDefaultGcsBucketNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetDefaultGcsBucketNameResponse.Size(m)
}
func (m *GetDefaultGcsBucketNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefaultGcsBucketNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefaultGcsBucketNameResponse proto.InternalMessageInfo

func (m *GetDefaultGcsBucketNameResponse) GetDefaultGcsBucketName() string {
	if m != nil && m.DefaultGcsBucketName != nil {
		return *m.DefaultGcsBucketName
	}
	return ""
}

func init() {
	proto.RegisterType((*AppIdentityServiceError)(nil), "appengine.AppIdentityServiceError")
	proto.RegisterType((*SignForAppRequest)(nil), "appengine.SignForAppRequest")
	proto.RegisterType((*SignForAppResponse)(nil), "appengine.SignForAppResponse")
	proto.RegisterType((*GetPublicCertificateForAppRequest)(nil), "appengine.GetPublicCertificateForAppRequest")
	proto.RegisterType((*PublicCertificate)(nil), "appengine.PublicCertificate")
	proto.RegisterType((*GetPublicCertificateForAppResponse)(nil), "appengine.GetPublicCertificateForAppResponse")
	proto.RegisterType((*GetServiceAccountNameRequest)(nil), "appengine.GetServiceAccountNameRequest")
	proto.RegisterType((*GetServiceAccountNameResponse)(nil), "appengine.GetServiceAccountNameResponse")
	proto.RegisterType((*GetAccessTokenRequest)(nil), "appengine.GetAccessTokenRequest")
	proto.RegisterType((*GetAccessTokenResponse)(nil), "appengine.GetAccessTokenResponse")
	proto.RegisterType((*GetDefaultGcsBucketNameRequest)(nil), "appengine.GetDefaultGcsBucketNameRequest")
	proto.RegisterType((*GetDefaultGcsBucketNameResponse)(nil), "appengine.GetDefaultGcsBucketNameResponse")
}

func init() {
	proto.RegisterFile("github.com/sgtsquiggs/appengine/internal/app_identity/app_identity_service.proto", fileDescriptor_app_identity_service_08a6e3f74b04cfa4)
}

var fileDescriptor_app_identity_service_08a6e3f74b04cfa4 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x6e, 0xda, 0x58,
	0x14, 0x1d, 0x26, 0x1a, 0x31, 0x6c, 0x12, 0x62, 0xce, 0x90, 0xcb, 0x8c, 0x32, 0xb9, 0x78, 0x1e,
	0x26, 0x0f, 0x15, 0x89, 0x2a, 0x45, 0x55, 0x1f, 0x8d, 0xed, 0x22, 0x54, 0x07, 0x53, 0x43, 0x9a,
	0xa8, 0x2f, 0xa7, 0xce, 0x61, 0xc7, 0x3d, 0x02, 0x9f, 0xe3, 0xda, 0x87, 0x0a, 0x3e, 0xa2, 0x3f,
	0xd2, 0x9f, 0xe8, 0x5b, 0xbf, 0xa5, 0x17, 0xb5, 0xdf, 0x50, 0xd9, 0x38, 0x5c, 0x92, 0x92, 0x37,
	0xbc, 0xf6, 0x5a, 0xcb, 0x6b, 0x2f, 0x6d, 0x0c, 0x4e, 0x20, 0x65, 0x30, 0xc4, 0x7a, 0x20, 0x87,
	0xbe, 0x08, 0xea, 0x32, 0x0e, 0x4e, 0xfc, 0x28, 0x42, 0x11, 0x70, 0x81, 0x27, 0x5c, 0x28, 0x8c,
	0x85, 0x3f, 0x4c, 0x21, 0xca, 0xfb, 0x28, 0x14, 0x57, 0x93, 0xa5, 0x07, 0x9a, 0x60, 0xfc, 0x8e,
	0x33, 0xac, 0x47, 0xb1, 0x54, 0x92, 0x94, 0x66, 0x5a, 0xfd, 0x53, 0x01, 0x76, 0x8c, 0x28, 0x6a,
	0xe5, 0xc4, 0xee, 0x94, 0x67, 0xc7, 0xb1, 0x8c, 0xf5, 0x0f, 0x05, 0x28, 0x65, 0xbf, 0x4c, 0xd9,
	0x47, 0x52, 0x86, 0x62, 0xf7, 0xc2, 0x34, 0xed, 0x6e, 0x57, 0xfb, 0x8d, 0x54, 0x61, 0xe3, 0xa2,
	0xfd, 0xbc, 0xed, 0x5e, 0xb6, 0x69, 0xd7, 0x74, 0x3b, 0xb6, 0x56, 0x22, 0x7f, 0x41, 0xa5, 0xe1,
	0xb8, 0x0d, 0xda, 0x73, 0x5d, 0xea, 0x18, 0x5e, 0xd3, 0xd6, 0x3e, 0x17, 0xc9, 0x36, 0x54, 0x2d,
	0xdb, 0xb0, 0x9c, 0x56, 0xdb, 0xa6, 0xf6, 0x95, 0x69, 0xdb, 0x96, 0x6d, 0x69, 0x5f, 0x8a, 0xa4,
	0x06, 0x9b, 0x6d, 0xb7, 0x47, 0x0d, 0xfa, 0xd2, 0x70, 0x5a, 0x16, 0x35, 0x3a, 0x1d, 0xed, 0x6b,
	0x91, 0x90, 0xb9, 0xab, 0xed, 0x79, 0xae, 0xa7, 0x7d, 0x2b, 0x12, 0x0d, 0xca, 0x19, 0xd3, 0x71,
	0xdc, 0x4b, 0xdb, 0xd2, 0xbe, 0xcf, 0xb4, 0xad, 0xf3, 0x8e, 0x63, 0x9f, 0xdb, 0xed, 0x9e, 0x6d,
	0x69, 0x3f, 0x8a, 0xfa, 0x13, 0xa8, 0x76, 0x79, 0x20, 0x9e, 0xc9, 0xd8, 0x88, 0x22, 0x0f, 0xdf,
	0x8e, 0x30, 0x51, 0x44, 0x87, 0x8d, 0xeb, 0x89, 0xc2, 0x84, 0x2a, 0x49, 0x13, 0x1e, 0x88, 0xdd,
	0xc2, 0x61, 0xe1, 0x78, 0xdd, 0x2b, 0x67, 0x60, 0x4f, 0xa6, 0x02, 0xfd, 0x0a, 0xc8, 0xa2, 0x30,
	0x89, 0xa4, 0x48, 0x90, 0xfc, 0x0d, 0x7f, 0x0e, 0x70, 0x42, 0x85, 0x1f, 0x62, 0x26, 0x2a, 0x79,
	0xc5, 0x01, 0x4e, 0xda, 0x7e, 0x88, 0xe4, 0x7f, 0xd8, 0x4c, 0xbd, 0x7c, 0x35, 0x8a, 0x91, 0x66,
	0x4e, 0xbb, 0xbf, 0x67, 0xb6, 0x95, 0x19, 0xdc, 0x48, 0x51, 0xfd, 0x3f, 0x38, 0x6a, 0xa2, 0xea,
	0x8c, 0xae, 0x87, 0x9c, 0x99, 0x18, 0x2b, 0x7e, 0xc3, 0x99, 0xaf, 0x70, 0x29, 0xa2, 0xfe, 0x1a,
	0xaa, 0xf7, 0x18, 0x0f, 0xbd, 0xfd, 0x14, 0x6a, 0xe3, 0xb3, 0xd3, 0xa7, 0x94, 0xcd, 0xe9, 0x34,
	0xc2, 0x30, 0x8b, 0x50, 0xf2, 0x48, 0x3a, 0x5b, 0x70, 0xea, 0x60, 0xa8, 0x7f, 0x2c, 0x80, 0xfe,
	0x50, 0x8e, 0x7c, 0xe3, 0x1e, 0xec, 0x44, 0x19, 0x65, 0xc9, 0x7a, 0xc8, 0x13, 0xb5, 0x5b, 0x38,
	0x5c, 0x3b, 0x2e, 0x3f, 0xde, 0xab, 0xcf, 0xce, 0xa6, 0x7e, 0xcf, 0xcc, 0xdb, 0x8a, 0xee, 0x42,
	0x0e, 0x4f, 0x14, 0x31, 0xe1, 0x20, 0xf4, 0xc7, 0x94, 0x0d, 0x39, 0x0a, 0x45, 0x99, 0xcf, 0xde,
	0x20, 0x55, 0x3c, 0x44, 0xca, 0x05, 0x4d, 0x90, 0x49, 0xd1, 0xcf, 0x92, 0xaf, 0x79, 0xff, 0x84,
	0xfe, 0xd8, 0xcc, 0x58, 0x66, 0x4a, 0xea, 0xf1, 0x10, 0x5b, 0xa2, 0x9b, 0x31, 0xf4, 0x7d, 0xd8,
	0x6b, 0xa2, 0xca, 0x6f, 0xd3, 0x60, 0x4c, 0x8e, 0x84, 0x4a, 0xcb, 0xb8, 0xed, 0xf0, 0x05, 0xfc,
	0xbb, 0x62, 0x9e, 0xef, 0x76, 0x0a, 0xb5, 0xfc, 0x1f, 0x40, 0xfd, 0xe9, 0x78, 0xb1, 0x5b, 0x92,
	0xdc, 0x53, 0xea, 0xef, 0x0b, 0xb0, 0xd5, 0x44, 0x65, 0x30, 0x86, 0x49, 0xd2, 0x93, 0x03, 0x14,
	0xb7, 0x37, 0x55, 0x83, 0x3f, 0x12, 0x26, 0x23, 0xcc, 0x5a, 0x29, 0x79, 0xd3, 0x07, 0xf2, 0x08,
	0xc8, 0xdd, 0x37, 0xf0, 0xdb, 0xd5, 0xb4, 0x65, 0xff, 0x56, 0x7f, 0x65, 0x9e, 0xb5, 0x95, 0x79,
	0xfa, 0xb0, 0x7d, 0x37, 0x4e, 0xbe, 0xdb, 0x11, 0xac, 0xfb, 0x19, 0x4c, 0x55, 0x8a, 0xe7, 0x3b,
	0x95, 0xfd, 0x39, 0x35, 0xbd, 0x58, 0x1c, 0x47, 0x3c, 0xf6, 0x15, 0x97, 0x22, 0xab, 0x3f, 0x4f,
	0x56, 0x99, 0xc3, 0x69, 0xe1, 0xfa, 0x21, 0xec, 0x37, 0x51, 0x59, 0x78, 0xe3, 0x8f, 0x86, 0xaa,
	0xc9, 0x92, 0xc6, 0x88, 0x0d, 0x70, 0xa9, 0xea, 0x2b, 0x38, 0x58, 0xc9, 0xc8, 0x03, 0x9d, 0xc1,
	0x4e, 0x7f, 0x3a, 0xa7, 0x01, 0x4b, 0xe8, 0x75, 0xc6, 0x58, 0xec, 0xbb, 0xd6, 0xff, 0x85, 0xbc,
	0x51, 0x79, 0xb5, 0xbe, 0xf8, 0xc9, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x37, 0x4c, 0x56, 0x38,
	0xf3, 0x04, 0x00, 0x00,
}
