// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: sso/sso.proto

package ssov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsAdminRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID to validate.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	mi := &file_sso_sso_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *IsAdminRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsAdminResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsAdmin       bool                   `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"` // Indicates whether the user is an admin.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	mi := &file_sso_sso_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`       // Email of the user to register.
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // Password of the user to register.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_sso_sso_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID of the registered user.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_sso_sso_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`               // Email of the user to login.
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`         // Password of the user to login.
	AppId         int32                  `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` // ID of the app to login to.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_sso_sso_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Auth token of the logged in user.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_sso_sso_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Auth token of the user to logout.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_sso_sso_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates whether the logout was successful.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_sso_sso_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{7}
}

func (x *LogoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_sso_sso_proto protoreflect.FileDescriptor

const file_sso_sso_proto_rawDesc = "" +
	"\n" +
	"\rsso/sso.proto\x12\x04auth\")\n" +
	"\x0eIsAdminRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\",\n" +
	"\x0fIsAdminResponse\x12\x19\n" +
	"\bis_admin\x18\x01 \x01(\bR\aisAdmin\"C\n" +
	"\x0fRegisterRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"+\n" +
	"\x10RegisterResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"W\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x15\n" +
	"\x06app_id\x18\x03 \x01(\x05R\x05appId\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"%\n" +
	"\rLogoutRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"*\n" +
	"\x0eLogoutResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xe0\x01\n" +
	"\x04Auth\x129\n" +
	"\bRegister\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x120\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x126\n" +
	"\aIsAdmin\x12\x14.auth.IsAdminRequest\x1a\x15.auth.IsAdminResponse\x123\n" +
	"\x06Logout\x12\x13.auth.LogoutRequest\x1a\x14.auth.LogoutResponseB\x19Z\x17cyansnbrst.sso.v1;ssov1b\x06proto3"

var (
	file_sso_sso_proto_rawDescOnce sync.Once
	file_sso_sso_proto_rawDescData []byte
)

func file_sso_sso_proto_rawDescGZIP() []byte {
	file_sso_sso_proto_rawDescOnce.Do(func() {
		file_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)))
	})
	return file_sso_sso_proto_rawDescData
}

var file_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_sso_proto_goTypes = []any{
	(*IsAdminRequest)(nil),   // 0: auth.IsAdminRequest
	(*IsAdminResponse)(nil),  // 1: auth.IsAdminResponse
	(*RegisterRequest)(nil),  // 2: auth.RegisterRequest
	(*RegisterResponse)(nil), // 3: auth.RegisterResponse
	(*LoginRequest)(nil),     // 4: auth.LoginRequest
	(*LoginResponse)(nil),    // 5: auth.LoginResponse
	(*LogoutRequest)(nil),    // 6: auth.LogoutRequest
	(*LogoutResponse)(nil),   // 7: auth.LogoutResponse
}
var file_sso_sso_proto_depIdxs = []int32{
	2, // 0: auth.Auth.Register:input_type -> auth.RegisterRequest
	4, // 1: auth.Auth.Login:input_type -> auth.LoginRequest
	0, // 2: auth.Auth.IsAdmin:input_type -> auth.IsAdminRequest
	6, // 3: auth.Auth.Logout:input_type -> auth.LogoutRequest
	3, // 4: auth.Auth.Register:output_type -> auth.RegisterResponse
	5, // 5: auth.Auth.Login:output_type -> auth.LoginResponse
	1, // 6: auth.Auth.IsAdmin:output_type -> auth.IsAdminResponse
	7, // 7: auth.Auth.Logout:output_type -> auth.LogoutResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_sso_proto_init() }
func file_sso_sso_proto_init() {
	if File_sso_sso_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_sso_proto_rawDesc), len(file_sso_sso_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_sso_proto_goTypes,
		DependencyIndexes: file_sso_sso_proto_depIdxs,
		MessageInfos:      file_sso_sso_proto_msgTypes,
	}.Build()
	File_sso_sso_proto = out.File
	file_sso_sso_proto_goTypes = nil
	file_sso_sso_proto_depIdxs = nil
}
