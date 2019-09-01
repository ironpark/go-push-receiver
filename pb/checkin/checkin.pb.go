// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin.proto

package checkin_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A concrete name/value pair sent to the device's Gservices database.
type GservicesSetting struct {
	Name                 []byte   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value                []byte   `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GservicesSetting) Reset()         { *m = GservicesSetting{} }
func (m *GservicesSetting) String() string { return proto.CompactTextString(m) }
func (*GservicesSetting) ProtoMessage()    {}
func (*GservicesSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{0}
}

func (m *GservicesSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GservicesSetting.Unmarshal(m, b)
}
func (m *GservicesSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GservicesSetting.Marshal(b, m, deterministic)
}
func (m *GservicesSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GservicesSetting.Merge(m, src)
}
func (m *GservicesSetting) XXX_Size() int {
	return xxx_messageInfo_GservicesSetting.Size(m)
}
func (m *GservicesSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_GservicesSetting.DiscardUnknown(m)
}

var xxx_messageInfo_GservicesSetting proto.InternalMessageInfo

func (m *GservicesSetting) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GservicesSetting) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Devices send this every few hours to tell us how they're doing.
type AndroidCheckinRequest struct {
	// IMEI (used by GSM phones) is sent and stored as 15 decimal
	// digits; the 15th is a check digit.
	Imei *string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// MEID (used by CDMA phones) is sent and stored as 14 hexadecimal
	// digits (no check digit).
	Meid *string `protobuf:"bytes,10,opt,name=meid" json:"meid,omitempty"`
	// MAC address (used by non-phone devices).  12 hexadecimal digits;
	// no separators (eg "0016E6513AC2", not "00:16:E6:51:3A:C2").
	MacAddr []string `protobuf:"bytes,9,rep,name=mac_addr,json=macAddr" json:"mac_addr,omitempty"`
	// An array parallel to mac_addr, describing the type of interface.
	// Currently accepted values: "wifi", "ethernet", "bluetooth".  If
	// not present, "wifi" is assumed.
	MacAddrType []string `protobuf:"bytes,19,rep,name=mac_addr_type,json=macAddrType" json:"mac_addr_type,omitempty"`
	// Serial number (a manufacturer-defined unique hardware
	// identifier).  Alphanumeric, case-insensitive.
	SerialNumber *string `protobuf:"bytes,16,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	// Older CDMA networks use an ESN (8 hex digits) instead of an MEID.
	Esn       *string              `protobuf:"bytes,17,opt,name=esn" json:"esn,omitempty"`
	Id        *int64               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LoggingId *int64               `protobuf:"varint,7,opt,name=logging_id,json=loggingId" json:"logging_id,omitempty"`
	Digest    *string              `protobuf:"bytes,3,opt,name=digest" json:"digest,omitempty"`
	Locale    *string              `protobuf:"bytes,6,opt,name=locale" json:"locale,omitempty"`
	Checkin   *AndroidCheckinProto `protobuf:"bytes,4,req,name=checkin" json:"checkin,omitempty"`
	// DEPRECATED, see AndroidCheckinProto.requested_group
	DesiredBuild *string `protobuf:"bytes,5,opt,name=desired_build,json=desiredBuild" json:"desired_build,omitempty"`
	// Blob of data from the Market app to be passed to Market API server
	MarketCheckin *string `protobuf:"bytes,8,opt,name=market_checkin,json=marketCheckin" json:"market_checkin,omitempty"`
	// SID cookies of any google accounts stored on the phone.  Not logged.
	AccountCookie []string `protobuf:"bytes,11,rep,name=account_cookie,json=accountCookie" json:"account_cookie,omitempty"`
	// Time zone.  Not currently logged.
	TimeZone *string `protobuf:"bytes,12,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	// Security token used to validate the checkin request.
	// Required for android IDs issued to Froyo+ devices, not for legacy IDs.
	SecurityToken *uint64 `protobuf:"fixed64,13,opt,name=security_token,json=securityToken" json:"security_token,omitempty"`
	// Version of checkin protocol.
	//
	// There are currently two versions:
	//
	// - version field missing: android IDs are assigned based on
	//   hardware identifiers.  unsecured in the sense that you can
	//   "unregister" someone's phone by sending a registration request
	//   with their IMEI/MEID/MAC.
	//
	// - version=2: android IDs are assigned randomly.  The device is
	//   sent a security token that must be included in all future
	//   checkins for that android id.
	//
	// - version=3: same as version 2, but the 'fragment' field is
	//   provided, and the device understands incremental updates to the
	//   gservices table (ie, only returning the keys whose values have
	//   changed.)
	//
	// (version=1 was skipped to avoid confusion with the "missing"
	// version field that is effectively version 1.)
	Version *int32 `protobuf:"varint,14,opt,name=version" json:"version,omitempty"`
	// OTA certs accepted by device (base-64 SHA-1 of cert files).  Not
	// logged.
	OtaCert []string `protobuf:"bytes,15,rep,name=ota_cert,json=otaCert" json:"ota_cert,omitempty"`
	// A single CheckinTask on the device may lead to multiple checkin
	// requests if there is too much log data to upload in a single
	// request.  For version 3 and up, this field will be filled in with
	// the number of the request, starting with 0.
	Fragment *int32 `protobuf:"varint,20,opt,name=fragment" json:"fragment,omitempty"`
	// For devices supporting multiple users, the name of the current
	// profile (they all check in independently, just as if they were
	// multiple physical devices).  This may not be set, even if the
	// device is using multiuser.  (checkin.user_number should be set to
	// the ordinal of the user.)
	UserName *string `protobuf:"bytes,21,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// For devices supporting multiple user profiles, the serial number
	// for the user checking in.  Not logged.  May not be set, even if
	// the device supportes multiuser.  checkin.user_number is the
	// ordinal of the user (0, 1, 2, ...), which may be reused if users
	// are deleted and re-created.  user_serial_number is never reused
	// (unless the device is wiped).
	UserSerialNumber     *int32   `protobuf:"varint,22,opt,name=user_serial_number,json=userSerialNumber" json:"user_serial_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AndroidCheckinRequest) Reset()         { *m = AndroidCheckinRequest{} }
func (m *AndroidCheckinRequest) String() string { return proto.CompactTextString(m) }
func (*AndroidCheckinRequest) ProtoMessage()    {}
func (*AndroidCheckinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{1}
}

func (m *AndroidCheckinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndroidCheckinRequest.Unmarshal(m, b)
}
func (m *AndroidCheckinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndroidCheckinRequest.Marshal(b, m, deterministic)
}
func (m *AndroidCheckinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndroidCheckinRequest.Merge(m, src)
}
func (m *AndroidCheckinRequest) XXX_Size() int {
	return xxx_messageInfo_AndroidCheckinRequest.Size(m)
}
func (m *AndroidCheckinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AndroidCheckinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AndroidCheckinRequest proto.InternalMessageInfo

func (m *AndroidCheckinRequest) GetImei() string {
	if m != nil && m.Imei != nil {
		return *m.Imei
	}
	return ""
}

func (m *AndroidCheckinRequest) GetMeid() string {
	if m != nil && m.Meid != nil {
		return *m.Meid
	}
	return ""
}

func (m *AndroidCheckinRequest) GetMacAddr() []string {
	if m != nil {
		return m.MacAddr
	}
	return nil
}

func (m *AndroidCheckinRequest) GetMacAddrType() []string {
	if m != nil {
		return m.MacAddrType
	}
	return nil
}

func (m *AndroidCheckinRequest) GetSerialNumber() string {
	if m != nil && m.SerialNumber != nil {
		return *m.SerialNumber
	}
	return ""
}

func (m *AndroidCheckinRequest) GetEsn() string {
	if m != nil && m.Esn != nil {
		return *m.Esn
	}
	return ""
}

func (m *AndroidCheckinRequest) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AndroidCheckinRequest) GetLoggingId() int64 {
	if m != nil && m.LoggingId != nil {
		return *m.LoggingId
	}
	return 0
}

func (m *AndroidCheckinRequest) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

func (m *AndroidCheckinRequest) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *AndroidCheckinRequest) GetCheckin() *AndroidCheckinProto {
	if m != nil {
		return m.Checkin
	}
	return nil
}

func (m *AndroidCheckinRequest) GetDesiredBuild() string {
	if m != nil && m.DesiredBuild != nil {
		return *m.DesiredBuild
	}
	return ""
}

func (m *AndroidCheckinRequest) GetMarketCheckin() string {
	if m != nil && m.MarketCheckin != nil {
		return *m.MarketCheckin
	}
	return ""
}

func (m *AndroidCheckinRequest) GetAccountCookie() []string {
	if m != nil {
		return m.AccountCookie
	}
	return nil
}

func (m *AndroidCheckinRequest) GetTimeZone() string {
	if m != nil && m.TimeZone != nil {
		return *m.TimeZone
	}
	return ""
}

func (m *AndroidCheckinRequest) GetSecurityToken() uint64 {
	if m != nil && m.SecurityToken != nil {
		return *m.SecurityToken
	}
	return 0
}

func (m *AndroidCheckinRequest) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *AndroidCheckinRequest) GetOtaCert() []string {
	if m != nil {
		return m.OtaCert
	}
	return nil
}

func (m *AndroidCheckinRequest) GetFragment() int32 {
	if m != nil && m.Fragment != nil {
		return *m.Fragment
	}
	return 0
}

func (m *AndroidCheckinRequest) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *AndroidCheckinRequest) GetUserSerialNumber() int32 {
	if m != nil && m.UserSerialNumber != nil {
		return *m.UserSerialNumber
	}
	return 0
}

// The response to the device.
type AndroidCheckinResponse struct {
	StatsOk  *bool  `protobuf:"varint,1,req,name=stats_ok,json=statsOk" json:"stats_ok,omitempty"`
	TimeMsec *int64 `protobuf:"varint,3,opt,name=time_msec,json=timeMsec" json:"time_msec,omitempty"`
	// Provisioning is sent if the request included an obsolete digest.
	//
	// For version <= 2, 'digest' contains the digest that should be
	// sent back to the server on the next checkin, and 'setting'
	// contains the entire gservices table (which replaces the entire
	// current table on the device).
	//
	// for version >= 3, 'digest' will be absent.  If 'settings_diff'
	// is false, then 'setting' contains the entire table, as in version
	// 2.  If 'settings_diff' is true, then 'delete_setting' contains
	// the keys to delete, and 'setting' contains only keys to be added
	// or for which the value has changed.  All other keys in the
	// current table should be left untouched.  If 'settings_diff' is
	// absent, don't touch the existing gservices table.
	//
	Digest               *string             `protobuf:"bytes,4,opt,name=digest" json:"digest,omitempty"`
	SettingsDiff         *bool               `protobuf:"varint,9,opt,name=settings_diff,json=settingsDiff" json:"settings_diff,omitempty"`
	DeleteSetting        []string            `protobuf:"bytes,10,rep,name=delete_setting,json=deleteSetting" json:"delete_setting,omitempty"`
	Setting              []*GservicesSetting `protobuf:"bytes,5,rep,name=setting" json:"setting,omitempty"`
	MarketOk             *bool               `protobuf:"varint,6,opt,name=market_ok,json=marketOk" json:"market_ok,omitempty"`
	AndroidId            *uint64             `protobuf:"fixed64,7,opt,name=android_id,json=androidId" json:"android_id,omitempty"`
	SecurityToken        *uint64             `protobuf:"fixed64,8,opt,name=security_token,json=securityToken" json:"security_token,omitempty"`
	VersionInfo          *string             `protobuf:"bytes,11,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AndroidCheckinResponse) Reset()         { *m = AndroidCheckinResponse{} }
func (m *AndroidCheckinResponse) String() string { return proto.CompactTextString(m) }
func (*AndroidCheckinResponse) ProtoMessage()    {}
func (*AndroidCheckinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{2}
}

func (m *AndroidCheckinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndroidCheckinResponse.Unmarshal(m, b)
}
func (m *AndroidCheckinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndroidCheckinResponse.Marshal(b, m, deterministic)
}
func (m *AndroidCheckinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndroidCheckinResponse.Merge(m, src)
}
func (m *AndroidCheckinResponse) XXX_Size() int {
	return xxx_messageInfo_AndroidCheckinResponse.Size(m)
}
func (m *AndroidCheckinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AndroidCheckinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AndroidCheckinResponse proto.InternalMessageInfo

func (m *AndroidCheckinResponse) GetStatsOk() bool {
	if m != nil && m.StatsOk != nil {
		return *m.StatsOk
	}
	return false
}

func (m *AndroidCheckinResponse) GetTimeMsec() int64 {
	if m != nil && m.TimeMsec != nil {
		return *m.TimeMsec
	}
	return 0
}

func (m *AndroidCheckinResponse) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

func (m *AndroidCheckinResponse) GetSettingsDiff() bool {
	if m != nil && m.SettingsDiff != nil {
		return *m.SettingsDiff
	}
	return false
}

func (m *AndroidCheckinResponse) GetDeleteSetting() []string {
	if m != nil {
		return m.DeleteSetting
	}
	return nil
}

func (m *AndroidCheckinResponse) GetSetting() []*GservicesSetting {
	if m != nil {
		return m.Setting
	}
	return nil
}

func (m *AndroidCheckinResponse) GetMarketOk() bool {
	if m != nil && m.MarketOk != nil {
		return *m.MarketOk
	}
	return false
}

func (m *AndroidCheckinResponse) GetAndroidId() uint64 {
	if m != nil && m.AndroidId != nil {
		return *m.AndroidId
	}
	return 0
}

func (m *AndroidCheckinResponse) GetSecurityToken() uint64 {
	if m != nil && m.SecurityToken != nil {
		return *m.SecurityToken
	}
	return 0
}

func (m *AndroidCheckinResponse) GetVersionInfo() string {
	if m != nil && m.VersionInfo != nil {
		return *m.VersionInfo
	}
	return ""
}

func init() {
	proto.RegisterType((*GservicesSetting)(nil), "checkin_proto.GservicesSetting")
	proto.RegisterType((*AndroidCheckinRequest)(nil), "checkin_proto.AndroidCheckinRequest")
	proto.RegisterType((*AndroidCheckinResponse)(nil), "checkin_proto.AndroidCheckinResponse")
}

func init() { proto.RegisterFile("checkin.proto", fileDescriptor_072e71e6019dc001) }

var fileDescriptor_072e71e6019dc001 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x74, 0x6d, 0x13, 0xb7, 0x19, 0xc5, 0x6c, 0x93, 0x19, 0x42, 0x84, 0x22, 0xa4,
	0x1c, 0xd0, 0x0e, 0xbb, 0x21, 0xed, 0xb2, 0x0d, 0x09, 0x76, 0x60, 0x43, 0xd9, 0x4e, 0x5c, 0x2c,
	0x2f, 0x7e, 0x29, 0x56, 0x12, 0xbb, 0xd8, 0xce, 0xa4, 0xf1, 0xff, 0xf0, 0x27, 0x72, 0x47, 0xb6,
	0x93, 0x69, 0x9d, 0x76, 0xf3, 0xfb, 0x3d, 0xe7, 0x93, 0xfd, 0xbe, 0xcf, 0x41, 0x69, 0xf9, 0x0b,
	0xca, 0x5a, 0xc8, 0xa3, 0x8d, 0x56, 0x56, 0xe1, 0xa1, 0xa4, 0xbe, 0x3c, 0xdc, 0x67, 0x92, 0x6b,
	0x25, 0x38, 0xdd, 0xda, 0xb5, 0x3a, 0x41, 0xcb, 0xaf, 0x06, 0xf4, 0x9d, 0x28, 0xc1, 0x5c, 0x83,
	0xb5, 0x42, 0xae, 0x31, 0x46, 0x3b, 0x92, 0xb5, 0x40, 0x46, 0x59, 0x94, 0x2f, 0x0a, 0xbf, 0xc6,
	0x7b, 0x68, 0x72, 0xc7, 0x9a, 0x0e, 0x48, 0xe4, 0x61, 0x28, 0x56, 0x7f, 0x27, 0x68, 0xff, 0x34,
	0xe8, 0x9e, 0x07, 0xd9, 0x02, 0x7e, 0x77, 0x60, 0xac, 0xd3, 0x10, 0x2d, 0x08, 0x32, 0xca, 0x46,
	0x79, 0x52, 0xf8, 0xb5, 0x63, 0x2d, 0x08, 0x4e, 0x50, 0x60, 0x6e, 0x8d, 0x5f, 0xa3, 0xb8, 0x65,
	0x25, 0x65, 0x9c, 0x6b, 0x92, 0x64, 0xe3, 0x3c, 0x29, 0x66, 0x2d, 0x2b, 0x4f, 0x39, 0xd7, 0x78,
	0x85, 0xd2, 0xa1, 0x45, 0xed, 0xfd, 0x06, 0xc8, 0x2b, 0xdf, 0x9f, 0xf7, 0xfd, 0x9b, 0xfb, 0x0d,
	0xe0, 0x0f, 0x28, 0x35, 0xa0, 0x05, 0x6b, 0xa8, 0xec, 0xda, 0x5b, 0xd0, 0x64, 0xe9, 0xb5, 0x17,
	0x01, 0x5e, 0x7a, 0x86, 0x97, 0x68, 0x0c, 0x46, 0x92, 0x97, 0xbe, 0xe5, 0x96, 0x78, 0x17, 0x45,
	0x82, 0x93, 0x28, 0x1b, 0xe5, 0xe3, 0x22, 0x12, 0x1c, 0xbf, 0x45, 0xa8, 0x51, 0xeb, 0xb5, 0x90,
	0x6b, 0x2a, 0x38, 0x99, 0x79, 0x9e, 0xf4, 0xe4, 0x82, 0xe3, 0x03, 0x34, 0xe5, 0x62, 0x0d, 0xc6,
	0x92, 0xb1, 0xd7, 0xe8, 0x2b, 0xc7, 0x1b, 0x55, 0xb2, 0x06, 0xc8, 0x34, 0xf0, 0x50, 0xe1, 0x13,
	0x34, 0xeb, 0xa7, 0x4c, 0x76, 0xb2, 0x28, 0x9f, 0x1f, 0xaf, 0x8e, 0xb6, 0xcc, 0x38, 0xda, 0x9e,
	0xd9, 0x0f, 0xc7, 0x8a, 0xe1, 0x13, 0x77, 0x27, 0x0e, 0x46, 0x68, 0xe0, 0xf4, 0xb6, 0x13, 0x0d,
	0x27, 0x93, 0x70, 0xa7, 0x1e, 0x9e, 0x39, 0x86, 0x3f, 0xa2, 0xdd, 0x96, 0xe9, 0x1a, 0xec, 0xe0,
	0x27, 0x89, 0xfd, 0xae, 0x34, 0xd0, 0x5e, 0xd9, 0x6d, 0x63, 0x65, 0xa9, 0x3a, 0x69, 0x69, 0xa9,
	0x54, 0x2d, 0x80, 0xcc, 0xfd, 0x10, 0xd3, 0x9e, 0x9e, 0x7b, 0x88, 0xdf, 0xa0, 0xc4, 0x8a, 0x16,
	0xe8, 0x1f, 0x25, 0x81, 0x2c, 0xbc, 0x50, 0xec, 0xc0, 0x4f, 0x25, 0xc1, 0x69, 0x18, 0x28, 0x3b,
	0x2d, 0xec, 0x3d, 0xb5, 0xaa, 0x06, 0x49, 0xd2, 0x6c, 0x94, 0x4f, 0x8b, 0x74, 0xa0, 0x37, 0x0e,
	0x62, 0x82, 0x66, 0x77, 0xa0, 0x8d, 0x50, 0x92, 0xec, 0x66, 0xa3, 0x7c, 0x52, 0x0c, 0xa5, 0xf3,
	0x58, 0x59, 0x46, 0x4b, 0xd0, 0x96, 0xbc, 0x08, 0x1e, 0x2b, 0xcb, 0xce, 0x41, 0x5b, 0x7c, 0x88,
	0xe2, 0x4a, 0xb3, 0x75, 0x0b, 0xd2, 0x92, 0x3d, 0xff, 0xd5, 0x43, 0xed, 0x0e, 0xd5, 0x19, 0xd0,
	0xd4, 0x67, 0x71, 0x3f, 0x1c, 0xca, 0x81, 0x4b, 0x97, 0xc7, 0x4f, 0x08, 0xfb, 0xe6, 0xb6, 0xfb,
	0x07, 0x5e, 0x62, 0xe9, 0x3a, 0xd7, 0x8f, 0x12, 0xb0, 0xfa, 0x17, 0xa1, 0x83, 0xa7, 0x39, 0x35,
	0x1b, 0x25, 0x0d, 0xb8, 0xc3, 0x19, 0xcb, 0xac, 0xa1, 0xaa, 0xf6, 0x81, 0x8f, 0x8b, 0x99, 0xaf,
	0xaf, 0xea, 0x87, 0xa9, 0xb4, 0x06, 0x4a, 0xef, 0xfc, 0x38, 0x4c, 0xe5, 0xbb, 0x81, 0xf2, 0x51,
	0x26, 0x76, 0xb6, 0x32, 0xe1, 0x13, 0xe9, 0xdf, 0x91, 0xa1, 0x5c, 0x54, 0x15, 0x49, 0xb2, 0x51,
	0x1e, 0xbb, 0x44, 0x06, 0xf8, 0x45, 0x54, 0x95, 0x1b, 0x29, 0x87, 0x06, 0x2c, 0xd0, 0x1e, 0x13,
	0x14, 0x6c, 0x09, 0x74, 0x78, 0x88, 0x9f, 0xd1, 0x6c, 0xe8, 0x4f, 0xb2, 0x71, 0x3e, 0x3f, 0x7e,
	0xf7, 0x24, 0x47, 0x4f, 0x9f, 0x6e, 0x31, 0xec, 0x77, 0x67, 0xef, 0xf3, 0xa1, 0x6a, 0x9f, 0xce,
	0xb8, 0x88, 0x03, 0xb8, 0xaa, 0x5d, 0xdc, 0x87, 0xbf, 0x41, 0x1f, 0xf7, 0x69, 0x91, 0xf4, 0xe4,
	0x82, 0x3f, 0x63, 0x78, 0xfc, 0x9c, 0xe1, 0xef, 0xd1, 0xa2, 0x77, 0x98, 0x0a, 0x59, 0x29, 0x32,
	0xf7, 0x73, 0x98, 0xf7, 0xec, 0x42, 0x56, 0xea, 0x2c, 0xfa, 0x36, 0xfe, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0xdd, 0x76, 0xa2, 0x97, 0x04, 0x00, 0x00,
}
