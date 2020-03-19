// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouderrorreporting/v1beta1/common.proto

package clouderrorreporting

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Description of a group of similar error events.
type ErrorGroup struct {
	// The group resource name.
	// Example: <code>projects/my-project-123/groups/my-groupid</code>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Group IDs are unique for a given project. If the same kind of error
	// occurs in different service contexts, it will receive the same group ID.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Associated tracking issues.
	TrackingIssues       []*TrackingIssue `protobuf:"bytes,3,rep,name=tracking_issues,json=trackingIssues,proto3" json:"tracking_issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ErrorGroup) Reset()         { *m = ErrorGroup{} }
func (m *ErrorGroup) String() string { return proto.CompactTextString(m) }
func (*ErrorGroup) ProtoMessage()    {}
func (*ErrorGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{0}
}

func (m *ErrorGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorGroup.Unmarshal(m, b)
}
func (m *ErrorGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorGroup.Marshal(b, m, deterministic)
}
func (m *ErrorGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorGroup.Merge(m, src)
}
func (m *ErrorGroup) XXX_Size() int {
	return xxx_messageInfo_ErrorGroup.Size(m)
}
func (m *ErrorGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorGroup proto.InternalMessageInfo

func (m *ErrorGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ErrorGroup) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *ErrorGroup) GetTrackingIssues() []*TrackingIssue {
	if m != nil {
		return m.TrackingIssues
	}
	return nil
}

// Information related to tracking the progress on resolving the error.
type TrackingIssue struct {
	// A URL pointing to a related entry in an issue tracking system.
	// Example: https://github.com/user/project/issues/4
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingIssue) Reset()         { *m = TrackingIssue{} }
func (m *TrackingIssue) String() string { return proto.CompactTextString(m) }
func (*TrackingIssue) ProtoMessage()    {}
func (*TrackingIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{1}
}

func (m *TrackingIssue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingIssue.Unmarshal(m, b)
}
func (m *TrackingIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingIssue.Marshal(b, m, deterministic)
}
func (m *TrackingIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingIssue.Merge(m, src)
}
func (m *TrackingIssue) XXX_Size() int {
	return xxx_messageInfo_TrackingIssue.Size(m)
}
func (m *TrackingIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingIssue.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingIssue proto.InternalMessageInfo

func (m *TrackingIssue) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// An error event which is returned by the Error Reporting system.
type ErrorEvent struct {
	// Time when the event occurred as provided in the error report.
	// If the report did not contain a timestamp, the time the error was received
	// by the Error Reporting system is used.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The `ServiceContext` for which this error was reported.
	ServiceContext *ServiceContext `protobuf:"bytes,2,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// The stack trace that was reported or logged by the service.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Data about the context in which the error occurred.
	Context              *ErrorContext `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ErrorEvent) Reset()         { *m = ErrorEvent{} }
func (m *ErrorEvent) String() string { return proto.CompactTextString(m) }
func (*ErrorEvent) ProtoMessage()    {}
func (*ErrorEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{2}
}

func (m *ErrorEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorEvent.Unmarshal(m, b)
}
func (m *ErrorEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorEvent.Marshal(b, m, deterministic)
}
func (m *ErrorEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorEvent.Merge(m, src)
}
func (m *ErrorEvent) XXX_Size() int {
	return xxx_messageInfo_ErrorEvent.Size(m)
}
func (m *ErrorEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorEvent proto.InternalMessageInfo

func (m *ErrorEvent) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *ErrorEvent) GetServiceContext() *ServiceContext {
	if m != nil {
		return m.ServiceContext
	}
	return nil
}

func (m *ErrorEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorEvent) GetContext() *ErrorContext {
	if m != nil {
		return m.Context
	}
	return nil
}

// Describes a running service that sends errors.
// Its version changes over time and multiple versions can run in parallel.
type ServiceContext struct {
	// An identifier of the service, such as the name of the
	// executable, job, or Google App Engine service name. This field is expected
	// to have a low number of values that are relatively stable over time, as
	// opposed to `version`, which can be changed whenever new code is deployed.
	//
	// Contains the service name for error reports extracted from Google
	// App Engine logs or `default` if the App Engine default service is used.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Represents the source code version that the developer provided,
	// which could represent a version label or a Git SHA-1 hash, for example.
	// For App Engine standard environment, the version is set to the version of
	// the app.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Type of the MonitoredResource. List of possible values:
	// https://cloud.google.com/monitoring/api/resources
	//
	// Value is set automatically for incoming errors and must not be set when
	// reporting errors.
	ResourceType         string   `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceContext) Reset()         { *m = ServiceContext{} }
func (m *ServiceContext) String() string { return proto.CompactTextString(m) }
func (*ServiceContext) ProtoMessage()    {}
func (*ServiceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{3}
}

func (m *ServiceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceContext.Unmarshal(m, b)
}
func (m *ServiceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceContext.Marshal(b, m, deterministic)
}
func (m *ServiceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceContext.Merge(m, src)
}
func (m *ServiceContext) XXX_Size() int {
	return xxx_messageInfo_ServiceContext.Size(m)
}
func (m *ServiceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceContext.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceContext proto.InternalMessageInfo

func (m *ServiceContext) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceContext) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceContext) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

// A description of the context in which an error occurred.
// This data should be provided by the application when reporting an error,
// unless the
// error report has been generated automatically from Google App Engine logs.
type ErrorContext struct {
	// The HTTP request which was processed when the error was
	// triggered.
	HttpRequest *HttpRequestContext `protobuf:"bytes,1,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// The user who caused or was affected by the crash.
	// This can be a user ID, an email address, or an arbitrary token that
	// uniquely identifies the user.
	// When sending an error report, leave this field empty if the user was not
	// logged in. In this case the
	// Error Reporting system will use other data, such as remote IP address, to
	// distinguish affected users. See `affected_users_count` in
	// `ErrorGroupStats`.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// The location in the source code where the decision was made to
	// report the error, usually the place where it was logged.
	// For a logged exception this would be the source line where the
	// exception is logged, usually close to the place where it was
	// caught.
	ReportLocation       *SourceLocation `protobuf:"bytes,3,opt,name=report_location,json=reportLocation,proto3" json:"report_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ErrorContext) Reset()         { *m = ErrorContext{} }
func (m *ErrorContext) String() string { return proto.CompactTextString(m) }
func (*ErrorContext) ProtoMessage()    {}
func (*ErrorContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{4}
}

func (m *ErrorContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorContext.Unmarshal(m, b)
}
func (m *ErrorContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorContext.Marshal(b, m, deterministic)
}
func (m *ErrorContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorContext.Merge(m, src)
}
func (m *ErrorContext) XXX_Size() int {
	return xxx_messageInfo_ErrorContext.Size(m)
}
func (m *ErrorContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorContext.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorContext proto.InternalMessageInfo

func (m *ErrorContext) GetHttpRequest() *HttpRequestContext {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *ErrorContext) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ErrorContext) GetReportLocation() *SourceLocation {
	if m != nil {
		return m.ReportLocation
	}
	return nil
}

// HTTP request data that is related to a reported error.
// This data should be provided by the application when reporting an error,
// unless the
// error report has been generated automatically from Google App Engine logs.
type HttpRequestContext struct {
	// The type of HTTP request, such as `GET`, `POST`, etc.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// The URL of the request.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// The user agent information that is provided with the request.
	UserAgent string `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The referrer information that is provided with the request.
	Referrer string `protobuf:"bytes,4,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// The HTTP response status code for the request.
	ResponseStatusCode int32 `protobuf:"varint,5,opt,name=response_status_code,json=responseStatusCode,proto3" json:"response_status_code,omitempty"`
	// The IP address from which the request originated.
	// This can be IPv4, IPv6, or a token which is derived from the
	// IP address, depending on the data that has been provided
	// in the error report.
	RemoteIp             string   `protobuf:"bytes,6,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpRequestContext) Reset()         { *m = HttpRequestContext{} }
func (m *HttpRequestContext) String() string { return proto.CompactTextString(m) }
func (*HttpRequestContext) ProtoMessage()    {}
func (*HttpRequestContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{5}
}

func (m *HttpRequestContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRequestContext.Unmarshal(m, b)
}
func (m *HttpRequestContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRequestContext.Marshal(b, m, deterministic)
}
func (m *HttpRequestContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRequestContext.Merge(m, src)
}
func (m *HttpRequestContext) XXX_Size() int {
	return xxx_messageInfo_HttpRequestContext.Size(m)
}
func (m *HttpRequestContext) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRequestContext.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRequestContext proto.InternalMessageInfo

func (m *HttpRequestContext) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpRequestContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpRequestContext) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HttpRequestContext) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *HttpRequestContext) GetResponseStatusCode() int32 {
	if m != nil {
		return m.ResponseStatusCode
	}
	return 0
}

func (m *HttpRequestContext) GetRemoteIp() string {
	if m != nil {
		return m.RemoteIp
	}
	return ""
}

// Indicates a location in the source code of the service for which errors are
// reported. `functionName` must be provided by the application when reporting
// an error, unless the error report contains a `message` with a supported
// exception stack trace. All fields are optional for the later case.
type SourceLocation struct {
	// The source code filename, which can include a truncated relative
	// path, or a full path from a production machine.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// 1-based. 0 indicates that the line number is unknown.
	LineNumber int32 `protobuf:"varint,2,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// Human-readable name of a function or method.
	// The value can include optional context like the class or package name.
	// For example, `my.package.MyClass.method` in case of Java.
	FunctionName         string   `protobuf:"bytes,4,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceLocation) Reset()         { *m = SourceLocation{} }
func (m *SourceLocation) String() string { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()    {}
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e89497b32fa506, []int{6}
}

func (m *SourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceLocation.Unmarshal(m, b)
}
func (m *SourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceLocation.Marshal(b, m, deterministic)
}
func (m *SourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceLocation.Merge(m, src)
}
func (m *SourceLocation) XXX_Size() int {
	return xxx_messageInfo_SourceLocation.Size(m)
}
func (m *SourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_SourceLocation proto.InternalMessageInfo

func (m *SourceLocation) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *SourceLocation) GetLineNumber() int32 {
	if m != nil {
		return m.LineNumber
	}
	return 0
}

func (m *SourceLocation) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorGroup)(nil), "google.devtools.clouderrorreporting.v1beta1.ErrorGroup")
	proto.RegisterType((*TrackingIssue)(nil), "google.devtools.clouderrorreporting.v1beta1.TrackingIssue")
	proto.RegisterType((*ErrorEvent)(nil), "google.devtools.clouderrorreporting.v1beta1.ErrorEvent")
	proto.RegisterType((*ServiceContext)(nil), "google.devtools.clouderrorreporting.v1beta1.ServiceContext")
	proto.RegisterType((*ErrorContext)(nil), "google.devtools.clouderrorreporting.v1beta1.ErrorContext")
	proto.RegisterType((*HttpRequestContext)(nil), "google.devtools.clouderrorreporting.v1beta1.HttpRequestContext")
	proto.RegisterType((*SourceLocation)(nil), "google.devtools.clouderrorreporting.v1beta1.SourceLocation")
}

func init() {
	proto.RegisterFile("google/devtools/clouderrorreporting/v1beta1/common.proto", fileDescriptor_d0e89497b32fa506)
}

var fileDescriptor_d0e89497b32fa506 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x4f, 0x1b, 0x39,
	0x14, 0xd7, 0x24, 0xfc, 0x8b, 0x03, 0x61, 0x65, 0xad, 0x56, 0x21, 0xbb, 0x2b, 0x20, 0x5c, 0x90,
	0x56, 0x3b, 0x53, 0xe8, 0xa5, 0xc0, 0xa1, 0x82, 0x08, 0x51, 0xa4, 0x0a, 0xa1, 0x09, 0xe5, 0x50,
	0xa1, 0x8e, 0x9c, 0x99, 0x97, 0xc9, 0xb4, 0x33, 0xb6, 0x6b, 0x7b, 0xa2, 0x22, 0xc4, 0x37, 0xe9,
	0x27, 0xe8, 0x27, 0xa9, 0xfa, 0x51, 0xaa, 0x4a, 0x3d, 0xb6, 0xc7, 0xca, 0x1e, 0x1b, 0x12, 0xc1,
	0xa1, 0x39, 0xd9, 0xbf, 0xf7, 0xff, 0xf7, 0xfc, 0xfc, 0xd0, 0xb3, 0x94, 0xb1, 0x34, 0x87, 0x20,
	0x81, 0xb1, 0x62, 0x2c, 0x97, 0x41, 0x9c, 0xb3, 0x32, 0x01, 0x21, 0x98, 0x10, 0xc0, 0x99, 0x50,
	0x19, 0x4d, 0x83, 0xf1, 0xce, 0x00, 0x14, 0xd9, 0x09, 0x62, 0x56, 0x14, 0x8c, 0xfa, 0x5c, 0x30,
	0xc5, 0xf0, 0x7f, 0x95, 0xa7, 0xef, 0x3c, 0xfd, 0x47, 0x3c, 0x7d, 0xeb, 0xd9, 0xf9, 0xc7, 0xa6,
	0x21, 0x3c, 0x0b, 0x08, 0xa5, 0x4c, 0x11, 0x95, 0x31, 0x2a, 0xab, 0x50, 0x9d, 0xb5, 0x09, 0xad,
	0x00, 0xc9, 0x4a, 0x11, 0x83, 0x55, 0xad, 0x5b, 0x95, 0x41, 0x83, 0x72, 0x18, 0xa8, 0xac, 0x00,
	0xa9, 0x48, 0xc1, 0x2b, 0x83, 0xee, 0x0f, 0x0f, 0xa1, 0x63, 0x9d, 0xf4, 0x44, 0xb0, 0x92, 0x63,
	0x8c, 0xe6, 0x28, 0x29, 0xa0, 0xed, 0x6d, 0x78, 0xdb, 0x8d, 0xd0, 0xdc, 0xf1, 0x1a, 0x5a, 0x4a,
	0xb5, 0x32, 0xca, 0x92, 0x76, 0xcd, 0xc8, 0x17, 0x0d, 0x3e, 0x4d, 0x70, 0x8c, 0x56, 0x95, 0x20,
	0xf1, 0xbb, 0x8c, 0xa6, 0x51, 0x26, 0x65, 0x09, 0xb2, 0x5d, 0xdf, 0xa8, 0x6f, 0x37, 0x77, 0xf7,
	0xfd, 0x19, 0xe8, 0xf9, 0x17, 0x36, 0xc6, 0xa9, 0x0e, 0x11, 0xb6, 0xd4, 0x24, 0x94, 0xfb, 0xaf,
	0xbe, 0x1e, 0x86, 0xe8, 0xff, 0xc7, 0xfc, 0xab, 0x1c, 0x84, 0x67, 0xd2, 0x8f, 0x59, 0x11, 0x4c,
	0xf0, 0xd8, 0xe4, 0x82, 0xbd, 0x85, 0x58, 0xc9, 0xe0, 0xc6, 0xde, 0x6e, 0x03, 0x53, 0xb4, 0x0c,
	0x6e, 0xcc, 0x79, 0xdb, 0xdd, 0x44, 0x2b, 0x53, 0x79, 0xf1, 0x1f, 0xa8, 0x5e, 0x8a, 0xdc, 0x52,
	0xd7, 0xd7, 0xee, 0xc7, 0x9a, 0x6d, 0xce, 0xf1, 0x18, 0xa8, 0xc2, 0x7b, 0x08, 0x81, 0xbe, 0x44,
	0xba, 0x89, 0xc6, 0xae, 0xb9, 0xdb, 0x71, 0x44, 0x5d, 0x87, 0xfd, 0x0b, 0xd7, 0xe1, 0xb0, 0x61,
	0xac, 0x35, 0xc6, 0x09, 0x5a, 0x95, 0x20, 0xc6, 0x59, 0x0c, 0x51, 0xcc, 0xa8, 0x82, 0x0f, 0xca,
	0xb4, 0xb2, 0xb9, 0x7b, 0x30, 0x53, 0xa3, 0xfa, 0x55, 0x8c, 0x5e, 0x15, 0x22, 0x6c, 0xc9, 0x29,
	0x8c, 0xdb, 0x68, 0xb1, 0x00, 0x29, 0x49, 0x0a, 0xed, 0x7a, 0xf5, 0x50, 0x16, 0xe2, 0x3e, 0x5a,
	0x74, 0x79, 0xe7, 0x4d, 0xde, 0xbd, 0x99, 0xf2, 0x9a, 0x26, 0xb8, 0xac, 0x2e, 0x52, 0x37, 0x43,
	0xad, 0xfe, 0x83, 0x02, 0x6c, 0x49, 0x6e, 0x52, 0x2c, 0xd4, 0x9a, 0x31, 0x08, 0x99, 0x31, 0xea,
	0x4a, 0xb3, 0x10, 0x6f, 0xa1, 0x15, 0x37, 0xb4, 0x91, 0xba, 0xe6, 0xd0, 0x9e, 0x33, 0xfa, 0x65,
	0x27, 0xbc, 0xb8, 0xe6, 0xd0, 0xfd, 0xe6, 0xa1, 0xe5, 0xc9, 0x22, 0xf0, 0x00, 0x2d, 0x8f, 0x94,
	0xe2, 0x91, 0x80, 0xf7, 0x25, 0x48, 0x65, 0x5f, 0xe3, 0xf9, 0x4c, 0xac, 0x5e, 0x28, 0xc5, 0xc3,
	0xca, 0xdf, 0x71, 0x6b, 0x8e, 0xee, 0x65, 0xfa, 0x33, 0x94, 0x12, 0x84, 0xa5, 0x62, 0xee, 0xfa,
	0x21, 0xab, 0x40, 0x51, 0xce, 0x62, 0xf3, 0x0b, 0x0d, 0x9f, 0x99, 0x1f, 0xd2, 0x50, 0x7b, 0x69,
	0x43, 0x84, 0xad, 0xca, 0xc2, 0xe1, 0xee, 0x67, 0x0f, 0xe1, 0x87, 0xd5, 0xe1, 0xbf, 0xd0, 0x42,
	0x01, 0x6a, 0xc4, 0x12, 0x3b, 0xa4, 0x16, 0xb9, 0xc9, 0xad, 0xdd, 0x4d, 0x2e, 0xfe, 0x17, 0x21,
	0x5d, 0x6e, 0x44, 0x52, 0xa0, 0xca, 0x76, 0xbc, 0xa1, 0x25, 0x87, 0x5a, 0x80, 0x3b, 0x68, 0x49,
	0xc0, 0x10, 0x84, 0x00, 0x61, 0xdb, 0x7d, 0x87, 0xf1, 0x13, 0xf4, 0xa7, 0x00, 0xc9, 0x19, 0x95,
	0x10, 0x49, 0x45, 0x54, 0x29, 0xa3, 0x98, 0x25, 0x60, 0xe6, 0x66, 0x3e, 0xc4, 0x4e, 0xd7, 0x37,
	0xaa, 0x1e, 0x4b, 0x00, 0xff, 0x8d, 0x1a, 0x02, 0x0a, 0xa6, 0x20, 0xca, 0x78, 0x7b, 0xc1, 0x85,
	0xd3, 0x82, 0x53, 0xde, 0x95, 0xa8, 0x35, 0x4d, 0x56, 0x9b, 0x0f, 0xb3, 0x1c, 0x22, 0x4e, 0xd4,
	0xc8, 0x12, 0x59, 0xd2, 0x82, 0x73, 0xa2, 0x46, 0x78, 0x1d, 0x35, 0xf3, 0x8c, 0x42, 0x44, 0xcb,
	0x62, 0x60, 0x5b, 0x3f, 0x1f, 0x22, 0x2d, 0x3a, 0x33, 0x12, 0x3d, 0x2e, 0xc3, 0x92, 0xc6, 0x3a,
	0x52, 0x64, 0x56, 0x95, 0x1d, 0x17, 0x27, 0x3c, 0x23, 0x05, 0x1c, 0x7d, 0xf7, 0x90, 0xde, 0xb6,
	0xb3, 0x3c, 0xc9, 0x51, 0xb3, 0x67, 0xd6, 0xf3, 0xb9, 0xfe, 0xc7, 0xe7, 0xde, 0xeb, 0x37, 0xd6,
	0x37, 0x65, 0x39, 0xa1, 0xa9, 0xcf, 0x44, 0x1a, 0xa4, 0x40, 0xcd, 0x2f, 0x0f, 0xee, 0xf7, 0xce,
	0x6f, 0x2d, 0xfe, 0x83, 0x47, 0x74, 0x3f, 0x3d, 0xef, 0x53, 0x6d, 0xeb, 0xa4, 0xca, 0xd1, 0xd3,
	0xfa, 0xea, 0x93, 0x85, 0x77, 0x75, 0x5d, 0xee, 0x1c, 0x69, 0xe7, 0x2f, 0xce, 0xea, 0xca, 0x58,
	0x5d, 0x4d, 0x5b, 0x5d, 0x5d, 0x56, 0x29, 0x06, 0x0b, 0xa6, 0xb2, 0xa7, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa9, 0x22, 0x13, 0x95, 0x91, 0x06, 0x00, 0x00,
}