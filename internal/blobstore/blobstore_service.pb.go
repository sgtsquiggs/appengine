// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/sgtsquiggs/appengine/internal/blobstore/blobstore_service.proto

package blobstore

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

type BlobstoreServiceError_ErrorCode int32

const (
	BlobstoreServiceError_OK                        BlobstoreServiceError_ErrorCode = 0
	BlobstoreServiceError_INTERNAL_ERROR            BlobstoreServiceError_ErrorCode = 1
	BlobstoreServiceError_URL_TOO_LONG              BlobstoreServiceError_ErrorCode = 2
	BlobstoreServiceError_PERMISSION_DENIED         BlobstoreServiceError_ErrorCode = 3
	BlobstoreServiceError_BLOB_NOT_FOUND            BlobstoreServiceError_ErrorCode = 4
	BlobstoreServiceError_DATA_INDEX_OUT_OF_RANGE   BlobstoreServiceError_ErrorCode = 5
	BlobstoreServiceError_BLOB_FETCH_SIZE_TOO_LARGE BlobstoreServiceError_ErrorCode = 6
	BlobstoreServiceError_ARGUMENT_OUT_OF_RANGE     BlobstoreServiceError_ErrorCode = 8
	BlobstoreServiceError_INVALID_BLOB_KEY          BlobstoreServiceError_ErrorCode = 9
)

var BlobstoreServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "URL_TOO_LONG",
	3: "PERMISSION_DENIED",
	4: "BLOB_NOT_FOUND",
	5: "DATA_INDEX_OUT_OF_RANGE",
	6: "BLOB_FETCH_SIZE_TOO_LARGE",
	8: "ARGUMENT_OUT_OF_RANGE",
	9: "INVALID_BLOB_KEY",
}
var BlobstoreServiceError_ErrorCode_value = map[string]int32{
	"OK":                        0,
	"INTERNAL_ERROR":            1,
	"URL_TOO_LONG":              2,
	"PERMISSION_DENIED":         3,
	"BLOB_NOT_FOUND":            4,
	"DATA_INDEX_OUT_OF_RANGE":   5,
	"BLOB_FETCH_SIZE_TOO_LARGE": 6,
	"ARGUMENT_OUT_OF_RANGE":     8,
	"INVALID_BLOB_KEY":          9,
}

func (x BlobstoreServiceError_ErrorCode) Enum() *BlobstoreServiceError_ErrorCode {
	p := new(BlobstoreServiceError_ErrorCode)
	*p = x
	return p
}
func (x BlobstoreServiceError_ErrorCode) String() string {
	return proto.EnumName(BlobstoreServiceError_ErrorCode_name, int32(x))
}
func (x *BlobstoreServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BlobstoreServiceError_ErrorCode_value, data, "BlobstoreServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = BlobstoreServiceError_ErrorCode(value)
	return nil
}
func (BlobstoreServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{0, 0}
}

type BlobstoreServiceError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlobstoreServiceError) Reset()         { *m = BlobstoreServiceError{} }
func (m *BlobstoreServiceError) String() string { return proto.CompactTextString(m) }
func (*BlobstoreServiceError) ProtoMessage()    {}
func (*BlobstoreServiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{0}
}
func (m *BlobstoreServiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlobstoreServiceError.Unmarshal(m, b)
}
func (m *BlobstoreServiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlobstoreServiceError.Marshal(b, m, deterministic)
}
func (dst *BlobstoreServiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlobstoreServiceError.Merge(dst, src)
}
func (m *BlobstoreServiceError) XXX_Size() int {
	return xxx_messageInfo_BlobstoreServiceError.Size(m)
}
func (m *BlobstoreServiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_BlobstoreServiceError.DiscardUnknown(m)
}

var xxx_messageInfo_BlobstoreServiceError proto.InternalMessageInfo

type CreateUploadURLRequest struct {
	SuccessPath               *string  `protobuf:"bytes,1,req,name=success_path,json=successPath" json:"success_path,omitempty"`
	MaxUploadSizeBytes        *int64   `protobuf:"varint,2,opt,name=max_upload_size_bytes,json=maxUploadSizeBytes" json:"max_upload_size_bytes,omitempty"`
	MaxUploadSizePerBlobBytes *int64   `protobuf:"varint,3,opt,name=max_upload_size_per_blob_bytes,json=maxUploadSizePerBlobBytes" json:"max_upload_size_per_blob_bytes,omitempty"`
	GsBucketName              *string  `protobuf:"bytes,4,opt,name=gs_bucket_name,json=gsBucketName" json:"gs_bucket_name,omitempty"`
	UrlExpiryTimeSeconds      *int32   `protobuf:"varint,5,opt,name=url_expiry_time_seconds,json=urlExpiryTimeSeconds" json:"url_expiry_time_seconds,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CreateUploadURLRequest) Reset()         { *m = CreateUploadURLRequest{} }
func (m *CreateUploadURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUploadURLRequest) ProtoMessage()    {}
func (*CreateUploadURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{1}
}
func (m *CreateUploadURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUploadURLRequest.Unmarshal(m, b)
}
func (m *CreateUploadURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUploadURLRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUploadURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUploadURLRequest.Merge(dst, src)
}
func (m *CreateUploadURLRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUploadURLRequest.Size(m)
}
func (m *CreateUploadURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUploadURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUploadURLRequest proto.InternalMessageInfo

func (m *CreateUploadURLRequest) GetSuccessPath() string {
	if m != nil && m.SuccessPath != nil {
		return *m.SuccessPath
	}
	return ""
}

func (m *CreateUploadURLRequest) GetMaxUploadSizeBytes() int64 {
	if m != nil && m.MaxUploadSizeBytes != nil {
		return *m.MaxUploadSizeBytes
	}
	return 0
}

func (m *CreateUploadURLRequest) GetMaxUploadSizePerBlobBytes() int64 {
	if m != nil && m.MaxUploadSizePerBlobBytes != nil {
		return *m.MaxUploadSizePerBlobBytes
	}
	return 0
}

func (m *CreateUploadURLRequest) GetGsBucketName() string {
	if m != nil && m.GsBucketName != nil {
		return *m.GsBucketName
	}
	return ""
}

func (m *CreateUploadURLRequest) GetUrlExpiryTimeSeconds() int32 {
	if m != nil && m.UrlExpiryTimeSeconds != nil {
		return *m.UrlExpiryTimeSeconds
	}
	return 0
}

type CreateUploadURLResponse struct {
	Url                  *string  `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUploadURLResponse) Reset()         { *m = CreateUploadURLResponse{} }
func (m *CreateUploadURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUploadURLResponse) ProtoMessage()    {}
func (*CreateUploadURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{2}
}
func (m *CreateUploadURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUploadURLResponse.Unmarshal(m, b)
}
func (m *CreateUploadURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUploadURLResponse.Marshal(b, m, deterministic)
}
func (dst *CreateUploadURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUploadURLResponse.Merge(dst, src)
}
func (m *CreateUploadURLResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUploadURLResponse.Size(m)
}
func (m *CreateUploadURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUploadURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUploadURLResponse proto.InternalMessageInfo

func (m *CreateUploadURLResponse) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type DeleteBlobRequest struct {
	BlobKey              []string `protobuf:"bytes,1,rep,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	Token                *string  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlobRequest) Reset()         { *m = DeleteBlobRequest{} }
func (m *DeleteBlobRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlobRequest) ProtoMessage()    {}
func (*DeleteBlobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{3}
}
func (m *DeleteBlobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlobRequest.Unmarshal(m, b)
}
func (m *DeleteBlobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlobRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBlobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlobRequest.Merge(dst, src)
}
func (m *DeleteBlobRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlobRequest.Size(m)
}
func (m *DeleteBlobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlobRequest proto.InternalMessageInfo

func (m *DeleteBlobRequest) GetBlobKey() []string {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

func (m *DeleteBlobRequest) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type FetchDataRequest struct {
	BlobKey              *string  `protobuf:"bytes,1,req,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	StartIndex           *int64   `protobuf:"varint,2,req,name=start_index,json=startIndex" json:"start_index,omitempty"`
	EndIndex             *int64   `protobuf:"varint,3,req,name=end_index,json=endIndex" json:"end_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchDataRequest) Reset()         { *m = FetchDataRequest{} }
func (m *FetchDataRequest) String() string { return proto.CompactTextString(m) }
func (*FetchDataRequest) ProtoMessage()    {}
func (*FetchDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{4}
}
func (m *FetchDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchDataRequest.Unmarshal(m, b)
}
func (m *FetchDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchDataRequest.Marshal(b, m, deterministic)
}
func (dst *FetchDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchDataRequest.Merge(dst, src)
}
func (m *FetchDataRequest) XXX_Size() int {
	return xxx_messageInfo_FetchDataRequest.Size(m)
}
func (m *FetchDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchDataRequest proto.InternalMessageInfo

func (m *FetchDataRequest) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func (m *FetchDataRequest) GetStartIndex() int64 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

func (m *FetchDataRequest) GetEndIndex() int64 {
	if m != nil && m.EndIndex != nil {
		return *m.EndIndex
	}
	return 0
}

type FetchDataResponse struct {
	Data                 []byte   `protobuf:"bytes,1000,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchDataResponse) Reset()         { *m = FetchDataResponse{} }
func (m *FetchDataResponse) String() string { return proto.CompactTextString(m) }
func (*FetchDataResponse) ProtoMessage()    {}
func (*FetchDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{5}
}
func (m *FetchDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchDataResponse.Unmarshal(m, b)
}
func (m *FetchDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchDataResponse.Marshal(b, m, deterministic)
}
func (dst *FetchDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchDataResponse.Merge(dst, src)
}
func (m *FetchDataResponse) XXX_Size() int {
	return xxx_messageInfo_FetchDataResponse.Size(m)
}
func (m *FetchDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchDataResponse proto.InternalMessageInfo

func (m *FetchDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloneBlobRequest struct {
	BlobKey              []byte   `protobuf:"bytes,1,req,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	MimeType             []byte   `protobuf:"bytes,2,req,name=mime_type,json=mimeType" json:"mime_type,omitempty"`
	TargetAppId          []byte   `protobuf:"bytes,3,req,name=target_app_id,json=targetAppId" json:"target_app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloneBlobRequest) Reset()         { *m = CloneBlobRequest{} }
func (m *CloneBlobRequest) String() string { return proto.CompactTextString(m) }
func (*CloneBlobRequest) ProtoMessage()    {}
func (*CloneBlobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{6}
}
func (m *CloneBlobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneBlobRequest.Unmarshal(m, b)
}
func (m *CloneBlobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneBlobRequest.Marshal(b, m, deterministic)
}
func (dst *CloneBlobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneBlobRequest.Merge(dst, src)
}
func (m *CloneBlobRequest) XXX_Size() int {
	return xxx_messageInfo_CloneBlobRequest.Size(m)
}
func (m *CloneBlobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneBlobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloneBlobRequest proto.InternalMessageInfo

func (m *CloneBlobRequest) GetBlobKey() []byte {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

func (m *CloneBlobRequest) GetMimeType() []byte {
	if m != nil {
		return m.MimeType
	}
	return nil
}

func (m *CloneBlobRequest) GetTargetAppId() []byte {
	if m != nil {
		return m.TargetAppId
	}
	return nil
}

type CloneBlobResponse struct {
	BlobKey              []byte   `protobuf:"bytes,1,req,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloneBlobResponse) Reset()         { *m = CloneBlobResponse{} }
func (m *CloneBlobResponse) String() string { return proto.CompactTextString(m) }
func (*CloneBlobResponse) ProtoMessage()    {}
func (*CloneBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{7}
}
func (m *CloneBlobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneBlobResponse.Unmarshal(m, b)
}
func (m *CloneBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneBlobResponse.Marshal(b, m, deterministic)
}
func (dst *CloneBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneBlobResponse.Merge(dst, src)
}
func (m *CloneBlobResponse) XXX_Size() int {
	return xxx_messageInfo_CloneBlobResponse.Size(m)
}
func (m *CloneBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloneBlobResponse proto.InternalMessageInfo

func (m *CloneBlobResponse) GetBlobKey() []byte {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

type DecodeBlobKeyRequest struct {
	BlobKey              []string `protobuf:"bytes,1,rep,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecodeBlobKeyRequest) Reset()         { *m = DecodeBlobKeyRequest{} }
func (m *DecodeBlobKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DecodeBlobKeyRequest) ProtoMessage()    {}
func (*DecodeBlobKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{8}
}
func (m *DecodeBlobKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeBlobKeyRequest.Unmarshal(m, b)
}
func (m *DecodeBlobKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeBlobKeyRequest.Marshal(b, m, deterministic)
}
func (dst *DecodeBlobKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeBlobKeyRequest.Merge(dst, src)
}
func (m *DecodeBlobKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DecodeBlobKeyRequest.Size(m)
}
func (m *DecodeBlobKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeBlobKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeBlobKeyRequest proto.InternalMessageInfo

func (m *DecodeBlobKeyRequest) GetBlobKey() []string {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

type DecodeBlobKeyResponse struct {
	Decoded              []string `protobuf:"bytes,1,rep,name=decoded" json:"decoded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecodeBlobKeyResponse) Reset()         { *m = DecodeBlobKeyResponse{} }
func (m *DecodeBlobKeyResponse) String() string { return proto.CompactTextString(m) }
func (*DecodeBlobKeyResponse) ProtoMessage()    {}
func (*DecodeBlobKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{9}
}
func (m *DecodeBlobKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeBlobKeyResponse.Unmarshal(m, b)
}
func (m *DecodeBlobKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeBlobKeyResponse.Marshal(b, m, deterministic)
}
func (dst *DecodeBlobKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeBlobKeyResponse.Merge(dst, src)
}
func (m *DecodeBlobKeyResponse) XXX_Size() int {
	return xxx_messageInfo_DecodeBlobKeyResponse.Size(m)
}
func (m *DecodeBlobKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeBlobKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeBlobKeyResponse proto.InternalMessageInfo

func (m *DecodeBlobKeyResponse) GetDecoded() []string {
	if m != nil {
		return m.Decoded
	}
	return nil
}

type CreateEncodedGoogleStorageKeyRequest struct {
	Filename             *string  `protobuf:"bytes,1,req,name=filename" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEncodedGoogleStorageKeyRequest) Reset()         { *m = CreateEncodedGoogleStorageKeyRequest{} }
func (m *CreateEncodedGoogleStorageKeyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEncodedGoogleStorageKeyRequest) ProtoMessage()    {}
func (*CreateEncodedGoogleStorageKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{10}
}
func (m *CreateEncodedGoogleStorageKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest.Unmarshal(m, b)
}
func (m *CreateEncodedGoogleStorageKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEncodedGoogleStorageKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest.Merge(dst, src)
}
func (m *CreateEncodedGoogleStorageKeyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest.Size(m)
}
func (m *CreateEncodedGoogleStorageKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEncodedGoogleStorageKeyRequest proto.InternalMessageInfo

func (m *CreateEncodedGoogleStorageKeyRequest) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

type CreateEncodedGoogleStorageKeyResponse struct {
	BlobKey              *string  `protobuf:"bytes,1,req,name=blob_key,json=blobKey" json:"blob_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEncodedGoogleStorageKeyResponse) Reset()         { *m = CreateEncodedGoogleStorageKeyResponse{} }
func (m *CreateEncodedGoogleStorageKeyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEncodedGoogleStorageKeyResponse) ProtoMessage()    {}
func (*CreateEncodedGoogleStorageKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blobstore_service_3604fb6033ea2e2e, []int{11}
}
func (m *CreateEncodedGoogleStorageKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse.Unmarshal(m, b)
}
func (m *CreateEncodedGoogleStorageKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse.Marshal(b, m, deterministic)
}
func (dst *CreateEncodedGoogleStorageKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse.Merge(dst, src)
}
func (m *CreateEncodedGoogleStorageKeyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse.Size(m)
}
func (m *CreateEncodedGoogleStorageKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEncodedGoogleStorageKeyResponse proto.InternalMessageInfo

func (m *CreateEncodedGoogleStorageKeyResponse) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func init() {
	proto.RegisterType((*BlobstoreServiceError)(nil), "appengine.BlobstoreServiceError")
	proto.RegisterType((*CreateUploadURLRequest)(nil), "appengine.CreateUploadURLRequest")
	proto.RegisterType((*CreateUploadURLResponse)(nil), "appengine.CreateUploadURLResponse")
	proto.RegisterType((*DeleteBlobRequest)(nil), "appengine.DeleteBlobRequest")
	proto.RegisterType((*FetchDataRequest)(nil), "appengine.FetchDataRequest")
	proto.RegisterType((*FetchDataResponse)(nil), "appengine.FetchDataResponse")
	proto.RegisterType((*CloneBlobRequest)(nil), "appengine.CloneBlobRequest")
	proto.RegisterType((*CloneBlobResponse)(nil), "appengine.CloneBlobResponse")
	proto.RegisterType((*DecodeBlobKeyRequest)(nil), "appengine.DecodeBlobKeyRequest")
	proto.RegisterType((*DecodeBlobKeyResponse)(nil), "appengine.DecodeBlobKeyResponse")
	proto.RegisterType((*CreateEncodedGoogleStorageKeyRequest)(nil), "appengine.CreateEncodedGoogleStorageKeyRequest")
	proto.RegisterType((*CreateEncodedGoogleStorageKeyResponse)(nil), "appengine.CreateEncodedGoogleStorageKeyResponse")
}

func init() {
	proto.RegisterFile("github.com/sgtsquiggs/appengine/internal/blobstore/blobstore_service.proto", fileDescriptor_blobstore_service_3604fb6033ea2e2e)
}

var fileDescriptor_blobstore_service_3604fb6033ea2e2e = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6e, 0xe3, 0x44,
	0x10, 0xc6, 0x4e, 0x7b, 0x8d, 0xa7, 0xe1, 0xe4, 0xae, 0x1a, 0x9a, 0x52, 0x01, 0xc1, 0x3a, 0xa4,
	0x48, 0xa0, 0x56, 0xfd, 0xc1, 0x03, 0xd8, 0xb5, 0x13, 0xac, 0xe6, 0xec, 0x6a, 0xe3, 0x20, 0xb8,
	0x3f, 0xab, 0x6d, 0x3c, 0xb8, 0x56, 0x1d, 0xaf, 0x59, 0x6f, 0x50, 0x73, 0x0f, 0xc1, 0xbb, 0xf1,
	0x16, 0x48, 0xbc, 0x04, 0xf2, 0xda, 0x6d, 0x73, 0x07, 0x77, 0xf7, 0x6f, 0xe7, 0xfb, 0xf6, 0x9b,
	0xf9, 0x66, 0x66, 0xb5, 0x30, 0xcd, 0x84, 0xc8, 0x0a, 0x3c, 0xcf, 0x44, 0xc1, 0xcb, 0xec, 0x5c,
	0xc8, 0xec, 0x82, 0x57, 0x15, 0x96, 0x59, 0x5e, 0xe2, 0x45, 0x5e, 0x2a, 0x94, 0x25, 0x2f, 0x2e,
	0x6e, 0x0b, 0x71, 0x5b, 0x2b, 0x21, 0xf1, 0xf9, 0xc4, 0x6a, 0x94, 0x7f, 0xe4, 0x2b, 0x3c, 0xaf,
	0xa4, 0x50, 0x82, 0x58, 0x4f, 0x2a, 0xe7, 0x1f, 0x03, 0x86, 0xde, 0xe3, 0xb5, 0x45, 0x7b, 0x2b,
	0x90, 0x52, 0x48, 0xe7, 0x2f, 0x03, 0x2c, 0x7d, 0xba, 0x12, 0x29, 0x92, 0x17, 0x60, 0xc6, 0xd7,
	0xf6, 0x67, 0x84, 0xc0, 0xcb, 0x30, 0x4a, 0x02, 0x1a, 0xb9, 0x73, 0x16, 0x50, 0x1a, 0x53, 0xdb,
	0x20, 0x36, 0x0c, 0x96, 0x74, 0xce, 0x92, 0x38, 0x66, 0xf3, 0x38, 0x9a, 0xd9, 0x26, 0x19, 0xc2,
	0xd1, 0x4d, 0x40, 0x5f, 0x87, 0x8b, 0x45, 0x18, 0x47, 0xcc, 0x0f, 0xa2, 0x30, 0xf0, 0xed, 0x5e,
	0x23, 0xf6, 0xe6, 0xb1, 0xc7, 0xa2, 0x38, 0x61, 0xd3, 0x78, 0x19, 0xf9, 0xf6, 0x1e, 0x39, 0x83,
	0x13, 0xdf, 0x4d, 0x5c, 0x16, 0x46, 0x7e, 0xf0, 0x0b, 0x8b, 0x97, 0x09, 0x8b, 0xa7, 0x8c, 0xba,
	0xd1, 0x2c, 0xb0, 0xf7, 0xc9, 0x57, 0x70, 0xaa, 0x05, 0xd3, 0x20, 0xb9, 0xfa, 0x89, 0x2d, 0xc2,
	0x37, 0x41, 0x5b, 0xc5, 0xa5, 0xb3, 0xc0, 0x7e, 0x41, 0x4e, 0x61, 0xe8, 0xd2, 0xd9, 0xf2, 0x75,
	0x10, 0x25, 0xef, 0x2a, 0xfb, 0xe4, 0x18, 0xec, 0x30, 0xfa, 0xd9, 0x9d, 0x87, 0x3e, 0xd3, 0x19,
	0xae, 0x83, 0x5f, 0x6d, 0xcb, 0xf9, 0xd3, 0x84, 0x2f, 0xae, 0x24, 0x72, 0x85, 0xcb, 0xaa, 0x10,
	0x3c, 0x5d, 0xd2, 0x39, 0xc5, 0xdf, 0x37, 0x58, 0x2b, 0xf2, 0x2d, 0x0c, 0xea, 0xcd, 0x6a, 0x85,
	0x75, 0xcd, 0x2a, 0xae, 0xee, 0x46, 0xc6, 0xd8, 0x9c, 0x58, 0xf4, 0xb0, 0xc3, 0x6e, 0xb8, 0xba,
	0x23, 0x97, 0x30, 0x5c, 0xf3, 0x07, 0xb6, 0xd1, 0x52, 0x56, 0xe7, 0x6f, 0x91, 0xdd, 0x6e, 0x15,
	0xd6, 0x23, 0x73, 0x6c, 0x4c, 0x7a, 0x94, 0xac, 0xf9, 0x43, 0x9b, 0x76, 0x91, 0xbf, 0x45, 0xaf,
	0x61, 0x88, 0x0b, 0x5f, 0xbf, 0x2f, 0xa9, 0x50, 0xb2, 0x66, 0x31, 0x9d, 0xb6, 0xa7, 0xb5, 0xa7,
	0xef, 0x68, 0x6f, 0x50, 0x36, 0x3b, 0x69, 0x53, 0xbc, 0x82, 0x97, 0x59, 0xcd, 0x6e, 0x37, 0xab,
	0x7b, 0x54, 0xac, 0xe4, 0x6b, 0x1c, 0xed, 0x8d, 0x8d, 0x89, 0x45, 0x07, 0x59, 0xed, 0x69, 0x30,
	0xe2, 0x6b, 0x24, 0x3f, 0xc2, 0xc9, 0x46, 0x16, 0x0c, 0x1f, 0xaa, 0x5c, 0x6e, 0x99, 0xca, 0xd7,
	0xcd, 0xce, 0x57, 0xa2, 0x4c, 0xeb, 0xd1, 0xfe, 0xd8, 0x98, 0xec, 0xd3, 0xe3, 0x8d, 0x2c, 0x02,
	0xcd, 0x26, 0xf9, 0x1a, 0x17, 0x2d, 0xe7, 0x7c, 0x0f, 0x27, 0xff, 0x99, 0x47, 0x5d, 0x89, 0xb2,
	0x46, 0x62, 0x43, 0x6f, 0x23, 0x8b, 0x6e, 0x0e, 0xcd, 0xd1, 0xf1, 0xe1, 0xc8, 0xc7, 0x02, 0x15,
	0x36, 0xe6, 0x1e, 0xe7, 0x76, 0x0a, 0x7d, 0xdd, 0xcd, 0x3d, 0x6e, 0x47, 0xc6, 0xb8, 0x37, 0xb1,
	0xe8, 0x41, 0x13, 0x5f, 0xe3, 0x96, 0x1c, 0xc3, 0xbe, 0x12, 0xf7, 0x58, 0xea, 0xf9, 0x58, 0xb4,
	0x0d, 0x9c, 0x7b, 0xb0, 0xa7, 0xa8, 0x56, 0x77, 0x3e, 0x57, 0xfc, 0xff, 0x93, 0x98, 0xbb, 0x49,
	0xbe, 0x81, 0xc3, 0x5a, 0x71, 0xa9, 0x58, 0x5e, 0xa6, 0xf8, 0x30, 0x32, 0xc7, 0xe6, 0xa4, 0x47,
	0x41, 0x43, 0x61, 0x83, 0x90, 0x33, 0xb0, 0xb0, 0x4c, 0x3b, 0xba, 0xa7, 0xe9, 0x3e, 0x96, 0xa9,
	0x26, 0x9d, 0x1f, 0xe0, 0x68, 0xa7, 0x58, 0xd7, 0xd9, 0x09, 0xec, 0xa5, 0x5c, 0xf1, 0xd1, 0xdf,
	0x07, 0x63, 0x73, 0x32, 0xf0, 0xcc, 0xbe, 0x41, 0x35, 0xe0, 0x94, 0x60, 0x5f, 0x15, 0xa2, 0xfc,
	0x48, 0x7f, 0xe6, 0x64, 0xf0, 0x6c, 0xed, 0x0c, 0xac, 0x75, 0x33, 0x68, 0xb5, 0xad, 0x50, 0x1b,
	0x1b, 0xd0, 0x7e, 0x03, 0x24, 0xdb, 0x0a, 0x89, 0x03, 0x9f, 0x2b, 0x2e, 0x33, 0x54, 0x8c, 0x57,
	0x15, 0xcb, 0x53, 0x6d, 0x6d, 0x40, 0x0f, 0x5b, 0xd0, 0xad, 0xaa, 0x30, 0x75, 0xce, 0xe1, 0x68,
	0xa7, 0x5e, 0xe7, 0xee, 0xc3, 0x05, 0x9d, 0x4b, 0x38, 0xf6, 0x71, 0x25, 0x52, 0x2d, 0xb8, 0xc6,
	0xed, 0xa7, 0x77, 0xe0, 0x5c, 0xc2, 0xf0, 0x3d, 0x49, 0x57, 0x66, 0x04, 0x07, 0xa9, 0x26, 0xd2,
	0x47, 0x49, 0x17, 0x3a, 0x1e, 0xbc, 0x6a, 0xdf, 0x44, 0x50, 0x6a, 0x60, 0xa6, 0x3f, 0x9d, 0x85,
	0x12, 0x92, 0x67, 0xb8, 0x53, 0xf5, 0x4b, 0xe8, 0xff, 0x96, 0x17, 0xa8, 0x9f, 0x64, 0xbb, 0xb4,
	0xa7, 0xd8, 0xf1, 0xe0, 0xbb, 0x4f, 0xe4, 0xf8, 0x40, 0xb7, 0xcf, 0xd6, 0xbd, 0xc3, 0x37, 0xd6,
	0xd3, 0x07, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xfb, 0x81, 0x94, 0xfb, 0x04, 0x00,
	0x00,
}
