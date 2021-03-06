// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/education_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Possible values for Education placeholder fields.
type EducationPlaceholderFieldEnum_EducationPlaceholderField int32

const (
	// Not specified.
	EducationPlaceholderFieldEnum_UNSPECIFIED EducationPlaceholderFieldEnum_EducationPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	EducationPlaceholderFieldEnum_UNKNOWN EducationPlaceholderFieldEnum_EducationPlaceholderField = 1
	// Data Type: STRING. Required. Combination of PROGRAM ID and LOCATION ID
	// must be unique per offer.
	EducationPlaceholderFieldEnum_PROGRAM_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 2
	// Data Type: STRING. Combination of PROGRAM ID and LOCATION ID must be
	// unique per offer.
	EducationPlaceholderFieldEnum_LOCATION_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with program name to be shown
	// in dynamic ad.
	EducationPlaceholderFieldEnum_PROGRAM_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 4
	// Data Type: STRING. Area of study that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_AREA_OF_STUDY EducationPlaceholderFieldEnum_EducationPlaceholderField = 5
	// Data Type: STRING. Description of program that can be shown in dynamic
	// ad.
	EducationPlaceholderFieldEnum_PROGRAM_DESCRIPTION EducationPlaceholderFieldEnum_EducationPlaceholderField = 6
	// Data Type: STRING. Name of school that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_SCHOOL_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 7
	// Data Type: STRING. Complete school address, including postal code.
	EducationPlaceholderFieldEnum_ADDRESS EducationPlaceholderFieldEnum_EducationPlaceholderField = 8
	// Data Type: URL. Image to be displayed in ads.
	EducationPlaceholderFieldEnum_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 9
	// Data Type: URL. Alternative hosted file of image to be used in the ad.
	EducationPlaceholderFieldEnum_ALTERNATIVE_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 10
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific program and its location).
	EducationPlaceholderFieldEnum_FINAL_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 11
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	EducationPlaceholderFieldEnum_FINAL_MOBILE_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 12
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	EducationPlaceholderFieldEnum_TRACKING_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	EducationPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 14
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	EducationPlaceholderFieldEnum_ANDROID_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 15
	// Data Type: STRING_LIST. List of recommended program IDs to show together
	// with this item.
	EducationPlaceholderFieldEnum_SIMILAR_PROGRAM_IDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 16
	// Data Type: STRING. iOS app link.
	EducationPlaceholderFieldEnum_IOS_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 17
	// Data Type: INT64. iOS app store ID.
	EducationPlaceholderFieldEnum_IOS_APP_STORE_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 18
)

var EducationPlaceholderFieldEnum_EducationPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PROGRAM_ID",
	3:  "LOCATION_ID",
	4:  "PROGRAM_NAME",
	5:  "AREA_OF_STUDY",
	6:  "PROGRAM_DESCRIPTION",
	7:  "SCHOOL_NAME",
	8:  "ADDRESS",
	9:  "THUMBNAIL_IMAGE_URL",
	10: "ALTERNATIVE_THUMBNAIL_IMAGE_URL",
	11: "FINAL_URLS",
	12: "FINAL_MOBILE_URLS",
	13: "TRACKING_URL",
	14: "CONTEXTUAL_KEYWORDS",
	15: "ANDROID_APP_LINK",
	16: "SIMILAR_PROGRAM_IDS",
	17: "IOS_APP_LINK",
	18: "IOS_APP_STORE_ID",
}
var EducationPlaceholderFieldEnum_EducationPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"PROGRAM_ID":                      2,
	"LOCATION_ID":                     3,
	"PROGRAM_NAME":                    4,
	"AREA_OF_STUDY":                   5,
	"PROGRAM_DESCRIPTION":             6,
	"SCHOOL_NAME":                     7,
	"ADDRESS":                         8,
	"THUMBNAIL_IMAGE_URL":             9,
	"ALTERNATIVE_THUMBNAIL_IMAGE_URL": 10,
	"FINAL_URLS":                      11,
	"FINAL_MOBILE_URLS":               12,
	"TRACKING_URL":                    13,
	"CONTEXTUAL_KEYWORDS":             14,
	"ANDROID_APP_LINK":                15,
	"SIMILAR_PROGRAM_IDS":             16,
	"IOS_APP_LINK":                    17,
	"IOS_APP_STORE_ID":                18,
}

func (x EducationPlaceholderFieldEnum_EducationPlaceholderField) String() string {
	return proto.EnumName(EducationPlaceholderFieldEnum_EducationPlaceholderField_name, int32(x))
}
func (EducationPlaceholderFieldEnum_EducationPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_education_placeholder_field_d137ac4ba72ea1ae, []int{0, 0}
}

// Values for Education placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type EducationPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EducationPlaceholderFieldEnum) Reset()         { *m = EducationPlaceholderFieldEnum{} }
func (m *EducationPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*EducationPlaceholderFieldEnum) ProtoMessage()    {}
func (*EducationPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_education_placeholder_field_d137ac4ba72ea1ae, []int{0}
}
func (m *EducationPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *EducationPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *EducationPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EducationPlaceholderFieldEnum.Merge(dst, src)
}
func (m *EducationPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_EducationPlaceholderFieldEnum.Size(m)
}
func (m *EducationPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_EducationPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_EducationPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EducationPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.EducationPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.EducationPlaceholderFieldEnum_EducationPlaceholderField", EducationPlaceholderFieldEnum_EducationPlaceholderField_name, EducationPlaceholderFieldEnum_EducationPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/education_placeholder_field.proto", fileDescriptor_education_placeholder_field_d137ac4ba72ea1ae)
}

var fileDescriptor_education_placeholder_field_d137ac4ba72ea1ae = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd9, 0x16, 0x76, 0xc1, 0xdd, 0x3f, 0xae, 0x01, 0x21, 0x0e, 0x05, 0xb1, 0xdc, 0xd3,
	0x4a, 0x1c, 0x39, 0xa0, 0x69, 0xe2, 0x76, 0xad, 0x3a, 0x76, 0x64, 0x27, 0x5d, 0x16, 0x55, 0xb2,
	0x42, 0x13, 0x42, 0xa5, 0x34, 0xa9, 0x9a, 0xed, 0xbe, 0x0d, 0x17, 0x8e, 0xbc, 0x04, 0x77, 0x2e,
	0xbc, 0x12, 0x72, 0x42, 0xb6, 0x17, 0xca, 0x25, 0x1a, 0xcf, 0x7c, 0xf3, 0x9b, 0xd8, 0xf3, 0xa1,
	0x0f, 0x59, 0x59, 0x66, 0x79, 0x3a, 0x8c, 0x93, 0x6a, 0xd8, 0x84, 0x36, 0xba, 0x1b, 0x0d, 0xd3,
	0x62, 0xb7, 0xae, 0x86, 0x69, 0xb2, 0x5b, 0xc6, 0xb7, 0xab, 0xb2, 0x30, 0x9b, 0x3c, 0x5e, 0xa6,
	0x5f, 0xcb, 0x3c, 0x49, 0xb7, 0xe6, 0xcb, 0x2a, 0xcd, 0x13, 0x67, 0xb3, 0x2d, 0x6f, 0x4b, 0x32,
	0x68, 0xba, 0x9c, 0x38, 0xa9, 0x9c, 0x7b, 0x80, 0x73, 0x37, 0x72, 0x6a, 0xc0, 0xe5, 0xcf, 0x2e,
	0x1a, 0xd0, 0x16, 0x12, 0xec, 0x19, 0x13, 0x8b, 0xa0, 0xc5, 0x6e, 0x7d, 0xf9, 0xad, 0x8b, 0x5e,
	0x1e, 0x54, 0x90, 0x0b, 0xd4, 0x8b, 0x84, 0x0e, 0xa8, 0xcb, 0x26, 0x8c, 0x7a, 0xf8, 0x01, 0xe9,
	0xa1, 0x93, 0x48, 0xcc, 0x84, 0xbc, 0x16, 0xf8, 0x88, 0x9c, 0x23, 0x14, 0x28, 0x39, 0x55, 0xe0,
	0x1b, 0xe6, 0xe1, 0x8e, 0x55, 0x73, 0xe9, 0x42, 0xc8, 0xa4, 0xb0, 0x89, 0x2e, 0xc1, 0xe8, 0xb4,
	0x15, 0x08, 0xf0, 0x29, 0x7e, 0x48, 0xfa, 0xe8, 0x0c, 0x14, 0x05, 0x23, 0x27, 0x46, 0x87, 0x91,
	0x77, 0x83, 0x1f, 0x91, 0x17, 0xe8, 0x69, 0x2b, 0xf2, 0xa8, 0x76, 0x15, 0x0b, 0x2c, 0x00, 0x1f,
	0x5b, 0x9c, 0x76, 0xaf, 0xa4, 0xe4, 0x4d, 0xf3, 0x89, 0x1d, 0x0e, 0x9e, 0xa7, 0xa8, 0xd6, 0xf8,
	0xb1, 0x6d, 0x0b, 0xaf, 0x22, 0x7f, 0x2c, 0x80, 0x71, 0xc3, 0x7c, 0x98, 0x52, 0x13, 0x29, 0x8e,
	0x9f, 0x90, 0xb7, 0xe8, 0x35, 0xf0, 0x90, 0x2a, 0x01, 0x21, 0x9b, 0x53, 0xf3, 0x2f, 0x11, 0xb2,
	0xbf, 0x3e, 0x61, 0x02, 0xb8, 0x3d, 0x6a, 0xdc, 0x23, 0xcf, 0x51, 0xbf, 0x39, 0xfb, 0x72, 0xcc,
	0x38, 0x6d, 0xd2, 0xa7, 0xf6, 0x02, 0xa1, 0x02, 0x77, 0xc6, 0xc4, 0xb4, 0x6e, 0x3c, 0xb3, 0x63,
	0x5d, 0x29, 0x42, 0xfa, 0x31, 0x8c, 0x80, 0x9b, 0x19, 0xbd, 0xb9, 0x96, 0xca, 0xd3, 0xf8, 0x9c,
	0x3c, 0x43, 0x18, 0x84, 0xa7, 0x24, 0xf3, 0x0c, 0x04, 0x81, 0xe1, 0x4c, 0xcc, 0xf0, 0x85, 0x95,
	0x6b, 0xe6, 0x33, 0x0e, 0xca, 0xec, 0x9f, 0x4a, 0x63, 0x6c, 0xc9, 0x4c, 0xea, 0xbd, 0xb4, 0x6f,
	0x01, 0x6d, 0x46, 0x87, 0x52, 0x51, 0xfb, 0x84, 0x64, 0xfc, 0xfb, 0x08, 0xbd, 0x59, 0x96, 0x6b,
	0xe7, 0xbf, 0x7b, 0x1e, 0xbf, 0x3a, 0xb8, 0xc2, 0xc0, 0xda, 0x24, 0x38, 0xfa, 0x34, 0xfe, 0x0b,
	0xc8, 0xca, 0x3c, 0x2e, 0x32, 0xa7, 0xdc, 0x66, 0xc3, 0x2c, 0x2d, 0x6a, 0x13, 0xb5, 0xce, 0xdb,
	0xac, 0xaa, 0x03, 0x46, 0x7c, 0x5f, 0x7f, 0xbf, 0x77, 0xba, 0x53, 0x80, 0x1f, 0x9d, 0xc1, 0xb4,
	0x41, 0x41, 0x52, 0x39, 0x4d, 0x68, 0xa3, 0xf9, 0xc8, 0xb1, 0x86, 0xaa, 0x7e, 0xb5, 0xf5, 0x05,
	0x24, 0xd5, 0xe2, 0xbe, 0xbe, 0x98, 0x8f, 0x16, 0x75, 0xfd, 0xf3, 0x71, 0x3d, 0xf4, 0xdd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x8b, 0x0c, 0x41, 0xfc, 0x02, 0x00, 0x00,
}
