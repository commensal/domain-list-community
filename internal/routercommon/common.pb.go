package routercommon

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

// Type of domain value.
type Domain_Type int32

const (
	// The value is used as is.
	Domain_Plain Domain_Type = 0
	// The value is used as a regular expression.
	Domain_Regex Domain_Type = 1
	// The value is a root domain.
	Domain_RootDomain Domain_Type = 2
	// The value is a domain.
	Domain_Full Domain_Type = 3
)

// Enum value maps for Domain_Type.
var (
	Domain_Type_name = map[int32]string{
		0: "Plain",
		1: "Regex",
		2: "RootDomain",
		3: "Full",
	}
	Domain_Type_value = map[string]int32{
		"Plain":      0,
		"Regex":      1,
		"RootDomain": 2,
		"Full":       3,
	}
)

func (x Domain_Type) Enum() *Domain_Type {
	p := new(Domain_Type)
	*p = x
	return p
}

func (x Domain_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Domain_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_app_router_routercommon_common_proto_enumTypes[0].Descriptor()
}

func (Domain_Type) Type() protoreflect.EnumType {
	return &file_app_router_routercommon_common_proto_enumTypes[0]
}

func (x Domain_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Domain_Type.Descriptor instead.
func (Domain_Type) EnumDescriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0, 0}
}

// Domain for routing decision.
type Domain struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Domain matching type.
	Type Domain_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.router.routercommon.Domain_Type" json:"type,omitempty"`
	// Domain value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Attributes of this domain. May be used for filtering.
	Attribute     []*Domain_Attribute `protobuf:"bytes,3,rep,name=attribute,proto3" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Domain) Reset() {
	*x = Domain{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetType() Domain_Type {
	if x != nil {
		return x.Type
	}
	return Domain_Plain
}

func (x *Domain) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Domain) GetAttribute() []*Domain_Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// IP for routing decision, in CIDR form.
type CIDR struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IP address, should be either 4 or 16 bytes.
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Number of leading ones in the network mask.
	Prefix        uint32 `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	IpAddr        string `protobuf:"bytes,68000,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CIDR) Reset() {
	*x = CIDR{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CIDR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIDR) ProtoMessage() {}

func (x *CIDR) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIDR.ProtoReflect.Descriptor instead.
func (*CIDR) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{1}
}

func (x *CIDR) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *CIDR) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

func (x *CIDR) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

type GeoIP struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	CountryCode  string                 `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Cidr         []*CIDR                `protobuf:"bytes,2,rep,name=cidr,proto3" json:"cidr,omitempty"`
	InverseMatch bool                   `protobuf:"varint,3,opt,name=inverse_match,json=inverseMatch,proto3" json:"inverse_match,omitempty"`
	// resource_hash instruct simplified config converter to load domain from geo file.
	ResourceHash  []byte `protobuf:"bytes,4,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash,omitempty"`
	Code          string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	FilePath      string `protobuf:"bytes,68000,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoIP) Reset() {
	*x = GeoIP{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIP) ProtoMessage() {}

func (x *GeoIP) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIP.ProtoReflect.Descriptor instead.
func (*GeoIP) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{2}
}

func (x *GeoIP) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GeoIP) GetCidr() []*CIDR {
	if x != nil {
		return x.Cidr
	}
	return nil
}

func (x *GeoIP) GetInverseMatch() bool {
	if x != nil {
		return x.InverseMatch
	}
	return false
}

func (x *GeoIP) GetResourceHash() []byte {
	if x != nil {
		return x.ResourceHash
	}
	return nil
}

func (x *GeoIP) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GeoIP) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GeoIPList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entry         []*GeoIP               `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoIPList) Reset() {
	*x = GeoIPList{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoIPList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIPList) ProtoMessage() {}

func (x *GeoIPList) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIPList.ProtoReflect.Descriptor instead.
func (*GeoIPList) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{3}
}

func (x *GeoIPList) GetEntry() []*GeoIP {
	if x != nil {
		return x.Entry
	}
	return nil
}

type GeoSite struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	CountryCode string                 `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Domain      []*Domain              `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	// resource_hash instruct simplified config converter to load domain from geo file.
	ResourceHash  []byte `protobuf:"bytes,3,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash,omitempty"`
	Code          string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	FilePath      string `protobuf:"bytes,68000,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoSite) Reset() {
	*x = GeoSite{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoSite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSite) ProtoMessage() {}

func (x *GeoSite) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoSite.ProtoReflect.Descriptor instead.
func (*GeoSite) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{4}
}

func (x *GeoSite) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GeoSite) GetDomain() []*Domain {
	if x != nil {
		return x.Domain
	}
	return nil
}

func (x *GeoSite) GetResourceHash() []byte {
	if x != nil {
		return x.ResourceHash
	}
	return nil
}

func (x *GeoSite) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GeoSite) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GeoSiteList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entry         []*GeoSite             `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoSiteList) Reset() {
	*x = GeoSiteList{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoSiteList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSiteList) ProtoMessage() {}

func (x *GeoSiteList) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoSiteList.ProtoReflect.Descriptor instead.
func (*GeoSiteList) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{5}
}

func (x *GeoSiteList) GetEntry() []*GeoSite {
	if x != nil {
		return x.Entry
	}
	return nil
}

type Domain_Attribute struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key   string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to TypedValue:
	//
	//	*Domain_Attribute_BoolValue
	//	*Domain_Attribute_IntValue
	TypedValue    isDomain_Attribute_TypedValue `protobuf_oneof:"typed_value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Domain_Attribute) Reset() {
	*x = Domain_Attribute{}
	mi := &file_app_router_routercommon_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Domain_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain_Attribute) ProtoMessage() {}

func (x *Domain_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain_Attribute.ProtoReflect.Descriptor instead.
func (*Domain_Attribute) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Domain_Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Domain_Attribute) GetTypedValue() isDomain_Attribute_TypedValue {
	if x != nil {
		return x.TypedValue
	}
	return nil
}

func (x *Domain_Attribute) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.TypedValue.(*Domain_Attribute_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

func (x *Domain_Attribute) GetIntValue() int64 {
	if x != nil {
		if x, ok := x.TypedValue.(*Domain_Attribute_IntValue); ok {
			return x.IntValue
		}
	}
	return 0
}

type isDomain_Attribute_TypedValue interface {
	isDomain_Attribute_TypedValue()
}

type Domain_Attribute_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Domain_Attribute_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

func (*Domain_Attribute_BoolValue) isDomain_Attribute_TypedValue() {}

func (*Domain_Attribute_IntValue) isDomain_Attribute_TypedValue() {}

var File_app_router_routercommon_common_proto protoreflect.FileDescriptor

const file_app_router_routercommon_common_proto_rawDesc = "" +
	"\x0a\x24\x61\x70\x70\x2f\x72\x6f\x75\x74\x65\x72\x2f\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2f\x63\x6f\x6d\x6d\x6f\x6e\x2e\x70\x72\x6f\x74\x6f\x12\x22" +
	"\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x22\xdd\x02\x0a\x06\x44" +
	"\x6f\x6d\x61\x69\x6e\x12\x43\x0a\x04\x74\x79\x70\x65\x18\x01\x20\x01\x28\x0e\x32\x2f\x2e\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75" +
	"\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2e\x44\x6f\x6d\x61\x69\x6e\x2e\x54\x79\x70\x65\x52\x04\x74\x79\x70\x65\x12\x14\x0a\x05\x76\x61" +
	"\x6c\x75\x65\x18\x02\x20\x01\x28\x09\x52\x05\x76\x61\x6c\x75\x65\x12\x52\x0a\x09\x61\x74\x74\x72\x69\x62\x75\x74\x65\x18\x03\x20\x03\x28\x0b\x32\x34\x2e\x76\x32" +
	"\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2e\x44\x6f\x6d\x61\x69\x6e\x2e" +
	"\x41\x74\x74\x72\x69\x62\x75\x74\x65\x52\x09\x61\x74\x74\x72\x69\x62\x75\x74\x65\x1a\x6c\x0a\x09\x41\x74\x74\x72\x69\x62\x75\x74\x65\x12\x10\x0a\x03\x6b\x65\x79" +
	"\x18\x01\x20\x01\x28\x09\x52\x03\x6b\x65\x79\x12\x1f\x0a\x0a\x62\x6f\x6f\x6c\x5f\x76\x61\x6c\x75\x65\x18\x02\x20\x01\x28\x08\x48\x00\x52\x09\x62\x6f\x6f\x6c\x56" +
	"\x61\x6c\x75\x65\x12\x1d\x0a\x09\x69\x6e\x74\x5f\x76\x61\x6c\x75\x65\x18\x03\x20\x01\x28\x03\x48\x00\x52\x08\x69\x6e\x74\x56\x61\x6c\x75\x65\x42\x0d\x0a\x0b\x74" +
	"\x79\x70\x65\x64\x5f\x76\x61\x6c\x75\x65\x22\x36\x0a\x04\x54\x79\x70\x65\x12\x09\x0a\x05\x50\x6c\x61\x69\x6e\x10\x00\x12\x09\x0a\x05\x52\x65\x67\x65\x78\x10\x01" +
	"\x12\x0e\x0a\x0a\x52\x6f\x6f\x74\x44\x6f\x6d\x61\x69\x6e\x10\x02\x12\x08\x0a\x04\x46\x75\x6c\x6c\x10\x03\x22\x49\x0a\x04\x43\x49\x44\x52\x12\x0e\x0a\x02\x69\x70" +
	"\x18\x01\x20\x01\x28\x0c\x52\x02\x69\x70\x12\x16\x0a\x06\x70\x72\x65\x66\x69\x78\x18\x02\x20\x01\x28\x0d\x52\x06\x70\x72\x65\x66\x69\x78\x12\x19\x0a\x07\x69\x70" +
	"\x5f\x61\x64\x64\x72\x18\xa0\x93\x04\x20\x01\x28\x09\x52\x06\x69\x70\x41\x64\x64\x72\x22\xe5\x01\x0a\x05\x47\x65\x6f\x49\x50\x12\x21\x0a\x0c\x63\x6f\x75\x6e\x74" +
	"\x72\x79\x5f\x63\x6f\x64\x65\x18\x01\x20\x01\x28\x09\x52\x0b\x63\x6f\x75\x6e\x74\x72\x79\x43\x6f\x64\x65\x12\x3c\x0a\x04\x63\x69\x64\x72\x18\x02\x20\x03\x28\x0b" +
	"\x32\x28\x2e\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2e\x43\x49" +
	"\x44\x52\x52\x04\x63\x69\x64\x72\x12\x23\x0a\x0d\x69\x6e\x76\x65\x72\x73\x65\x5f\x6d\x61\x74\x63\x68\x18\x03\x20\x01\x28\x08\x52\x0c\x69\x6e\x76\x65\x72\x73\x65" +
	"\x4d\x61\x74\x63\x68\x12\x23\x0a\x0d\x72\x65\x73\x6f\x75\x72\x63\x65\x5f\x68\x61\x73\x68\x18\x04\x20\x01\x28\x0c\x52\x0c\x72\x65\x73\x6f\x75\x72\x63\x65\x48\x61" +
	"\x73\x68\x12\x12\x0a\x04\x63\x6f\x64\x65\x18\x05\x20\x01\x28\x09\x52\x04\x63\x6f\x64\x65\x12\x1d\x0a\x09\x66\x69\x6c\x65\x5f\x70\x61\x74\x68\x18\xa0\x93\x04\x20" +
	"\x01\x28\x09\x52\x08\x66\x69\x6c\x65\x50\x61\x74\x68\x22\x4c\x0a\x09\x47\x65\x6f\x49\x50\x4c\x69\x73\x74\x12\x3f\x0a\x05\x65\x6e\x74\x72\x79\x18\x01\x20\x03\x28" +
	"\x0b\x32\x29\x2e\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2e\x47" +
	"\x65\x6f\x49\x50\x52\x05\x65\x6e\x74\x72\x79\x22\xc8\x01\x0a\x07\x47\x65\x6f\x53\x69\x74\x65\x12\x21\x0a\x0c\x63\x6f\x75\x6e\x74\x72\x79\x5f\x63\x6f\x64\x65\x18" +
	"\x01\x20\x01\x28\x09\x52\x0b\x63\x6f\x75\x6e\x74\x72\x79\x43\x6f\x64\x65\x12\x42\x0a\x06\x64\x6f\x6d\x61\x69\x6e\x18\x02\x20\x03\x28\x0b\x32\x2a\x2e\x76\x32\x72" +
	"\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x2e\x44\x6f\x6d\x61\x69\x6e\x52\x06" +
	"\x64\x6f\x6d\x61\x69\x6e\x12\x23\x0a\x0d\x72\x65\x73\x6f\x75\x72\x63\x65\x5f\x68\x61\x73\x68\x18\x03\x20\x01\x28\x0c\x52\x0c\x72\x65\x73\x6f\x75\x72\x63\x65\x48" +
	"\x61\x73\x68\x12\x12\x0a\x04\x63\x6f\x64\x65\x18\x04\x20\x01\x28\x09\x52\x04\x63\x6f\x64\x65\x12\x1d\x0a\x09\x66\x69\x6c\x65\x5f\x70\x61\x74\x68\x18\xa0\x93\x04" +
	"\x20\x01\x28\x09\x52\x08\x66\x69\x6c\x65\x50\x61\x74\x68\x22\x50\x0a\x0b\x47\x65\x6f\x53\x69\x74\x65\x4c\x69\x73\x74\x12\x41\x0a\x05\x65\x6e\x74\x72\x79\x18\x01" +
	"\x20\x03\x28\x0b\x32\x2b\x2e\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f" +
	"\x6e\x2e\x47\x65\x6f\x53\x69\x74\x65\x52\x05\x65\x6e\x74\x72\x79\x42\x8d\x01\x0a\x26\x63\x6f\x6d\x2e\x76\x32\x72\x61\x79\x2e\x63\x6f\x72\x65\x2e\x61\x70\x70\x2e" +
	"\x72\x6f\x75\x74\x65\x72\x2e\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e\x50\x01\x5a\x3c\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x2f\x76\x32\x66\x6c\x79\x2f" +
	"\x64\x6f\x6d\x61\x69\x6e\x2d\x6c\x69\x73\x74\x2d\x63\x6f\x6d\x6d\x75\x6e\x69\x74\x79\x2f\x69\x6e\x74\x65\x72\x6e\x61\x6c\x2f\x72\x6f\x75\x74\x65\x72\x63\x6f\x6d" +
	"\x6d\x6f\x6e\xaa\x02\x22\x56\x32\x52\x61\x79\x2e\x43\x6f\x72\x65\x2e\x41\x70\x70\x2e\x52\x6f\x75\x74\x65\x72\x2e\x52\x6f\x75\x74\x65\x72\x63\x6f\x6d\x6d\x6f\x6e" +
	"\x62\x06\x70\x72\x6f\x74\x6f\x33"
var (
	file_app_router_routercommon_common_proto_rawDescOnce sync.Once
	file_app_router_routercommon_common_proto_rawDescData []byte
)

func file_app_router_routercommon_common_proto_rawDescGZIP() []byte {
	file_app_router_routercommon_common_proto_rawDescOnce.Do(func() {
		file_app_router_routercommon_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_router_routercommon_common_proto_rawDesc), len(file_app_router_routercommon_common_proto_rawDesc)))
	})
	return file_app_router_routercommon_common_proto_rawDescData
}

var file_app_router_routercommon_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_router_routercommon_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_router_routercommon_common_proto_goTypes = []any{
	(Domain_Type)(0),         // 0: v2ray.core.app.router.routercommon.Domain.Type
	(*Domain)(nil),           // 1: v2ray.core.app.router.routercommon.Domain
	(*CIDR)(nil),             // 2: v2ray.core.app.router.routercommon.CIDR
	(*GeoIP)(nil),            // 3: v2ray.core.app.router.routercommon.GeoIP
	(*GeoIPList)(nil),        // 4: v2ray.core.app.router.routercommon.GeoIPList
	(*GeoSite)(nil),          // 5: v2ray.core.app.router.routercommon.GeoSite
	(*GeoSiteList)(nil),      // 6: v2ray.core.app.router.routercommon.GeoSiteList
	(*Domain_Attribute)(nil), // 7: v2ray.core.app.router.routercommon.Domain.Attribute
}
var file_app_router_routercommon_common_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.router.routercommon.Domain.type:type_name -> v2ray.core.app.router.routercommon.Domain.Type
	7, // 1: v2ray.core.app.router.routercommon.Domain.attribute:type_name -> v2ray.core.app.router.routercommon.Domain.Attribute
	2, // 2: v2ray.core.app.router.routercommon.GeoIP.cidr:type_name -> v2ray.core.app.router.routercommon.CIDR
	3, // 3: v2ray.core.app.router.routercommon.GeoIPList.entry:type_name -> v2ray.core.app.router.routercommon.GeoIP
	1, // 4: v2ray.core.app.router.routercommon.GeoSite.domain:type_name -> v2ray.core.app.router.routercommon.Domain
	5, // 5: v2ray.core.app.router.routercommon.GeoSiteList.entry:type_name -> v2ray.core.app.router.routercommon.GeoSite
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_app_router_routercommon_common_proto_init() }
func file_app_router_routercommon_common_proto_init() {
	if File_app_router_routercommon_common_proto != nil {
		return
	}
	file_app_router_routercommon_common_proto_msgTypes[6].OneofWrappers = []any{
		(*Domain_Attribute_BoolValue)(nil),
		(*Domain_Attribute_IntValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_router_routercommon_common_proto_rawDesc), len(file_app_router_routercommon_common_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_router_routercommon_common_proto_goTypes,
		DependencyIndexes: file_app_router_routercommon_common_proto_depIdxs,
		EnumInfos:         file_app_router_routercommon_common_proto_enumTypes,
		MessageInfos:      file_app_router_routercommon_common_proto_msgTypes,
	}.Build()
	File_app_router_routercommon_common_proto = out.File
	file_app_router_routercommon_common_proto_goTypes = nil
	file_app_router_routercommon_common_proto_depIdxs = nil
}
