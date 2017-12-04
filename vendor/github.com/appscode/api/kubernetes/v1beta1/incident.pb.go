// Code generated by protoc-gen-go. DO NOT EDIT.
// source: incident.proto

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Incident struct {
	Phid                 string               `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    *appscode_dtypes.Uid `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string               `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string               `protobuf:"bytes,4,opt,name=kubernetes_object_type,json=kubernetesObjectType" json:"kubernetes_object_type,omitempty"`
	KubernetesObjectName string               `protobuf:"bytes,5,opt,name=kubernetes_object_name,json=kubernetesObjectName" json:"kubernetes_object_name,omitempty"`
	KubernetesAlertName  string               `protobuf:"bytes,6,opt,name=kubernetes_alert_name,json=kubernetesAlertName" json:"kubernetes_alert_name,omitempty"`
	IcingaHost           string               `protobuf:"bytes,7,opt,name=icinga_host,json=icingaHost" json:"icinga_host,omitempty"`
	IcingaService        string               `protobuf:"bytes,8,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	Type                 string               `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	State                string               `protobuf:"bytes,10,opt,name=state" json:"state,omitempty"`
	User                 *appscode_dtypes.Uid `protobuf:"bytes,11,opt,name=user" json:"user,omitempty"`
	// Timestamp of first reported event
	ReportedAt int64 `protobuf:"varint,12,opt,name=reported_at,json=reportedAt" json:"reported_at,omitempty"`
	// Timestamp of first acknowledgement
	AcknowledgedAt int64             `protobuf:"varint,13,opt,name=acknowledged_at,json=acknowledgedAt" json:"acknowledged_at,omitempty"`
	RecoveredAt    int64             `protobuf:"varint,14,opt,name=recovered_at,json=recoveredAt" json:"recovered_at,omitempty"`
	IcingawebUrl   string            `protobuf:"bytes,15,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Events         []*Incident_Event `protobuf:"bytes,16,rep,name=events" json:"events,omitempty"`
}

func (m *Incident) Reset()                    { *m = Incident{} }
func (m *Incident) String() string            { return proto.CompactTextString(m) }
func (*Incident) ProtoMessage()               {}
func (*Incident) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Incident) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Incident) GetKubernetesCluster() *appscode_dtypes.Uid {
	if m != nil {
		return m.KubernetesCluster
	}
	return nil
}

func (m *Incident) GetKubernetesNamespace() string {
	if m != nil {
		return m.KubernetesNamespace
	}
	return ""
}

func (m *Incident) GetKubernetesObjectType() string {
	if m != nil {
		return m.KubernetesObjectType
	}
	return ""
}

func (m *Incident) GetKubernetesObjectName() string {
	if m != nil {
		return m.KubernetesObjectName
	}
	return ""
}

func (m *Incident) GetKubernetesAlertName() string {
	if m != nil {
		return m.KubernetesAlertName
	}
	return ""
}

func (m *Incident) GetIcingaHost() string {
	if m != nil {
		return m.IcingaHost
	}
	return ""
}

func (m *Incident) GetIcingaService() string {
	if m != nil {
		return m.IcingaService
	}
	return ""
}

func (m *Incident) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Incident) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Incident) GetUser() *appscode_dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Incident) GetReportedAt() int64 {
	if m != nil {
		return m.ReportedAt
	}
	return 0
}

func (m *Incident) GetAcknowledgedAt() int64 {
	if m != nil {
		return m.AcknowledgedAt
	}
	return 0
}

func (m *Incident) GetRecoveredAt() int64 {
	if m != nil {
		return m.RecoveredAt
	}
	return 0
}

func (m *Incident) GetIcingawebUrl() string {
	if m != nil {
		return m.IcingawebUrl
	}
	return ""
}

func (m *Incident) GetEvents() []*Incident_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Incident_Event struct {
	Type       string               `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	State      string               `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	ReportedAt int64                `protobuf:"varint,3,opt,name=reported_at,json=reportedAt" json:"reported_at,omitempty"`
	User       *appscode_dtypes.Uid `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Comment    string               `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
}

func (m *Incident_Event) Reset()                    { *m = Incident_Event{} }
func (m *Incident_Event) String() string            { return proto.CompactTextString(m) }
func (*Incident_Event) ProtoMessage()               {}
func (*Incident_Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Incident_Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Incident_Event) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Incident_Event) GetReportedAt() int64 {
	if m != nil {
		return m.ReportedAt
	}
	return 0
}

func (m *Incident_Event) GetUser() *appscode_dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Incident_Event) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

// Next Id: 6
type IncidentListRequest struct {
	KubernetesCluster    string   `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string   `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string   `protobuf:"bytes,3,opt,name=kubernetes_object_type,json=kubernetesObjectType" json:"kubernetes_object_type,omitempty"`
	KubernetesObjectName string   `protobuf:"bytes,4,opt,name=kubernetes_object_name,json=kubernetesObjectName" json:"kubernetes_object_name,omitempty"`
	States               []string `protobuf:"bytes,5,rep,name=states" json:"states,omitempty"`
}

func (m *IncidentListRequest) Reset()                    { *m = IncidentListRequest{} }
func (m *IncidentListRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentListRequest) ProtoMessage()               {}
func (*IncidentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *IncidentListRequest) GetKubernetesCluster() string {
	if m != nil {
		return m.KubernetesCluster
	}
	return ""
}

func (m *IncidentListRequest) GetKubernetesNamespace() string {
	if m != nil {
		return m.KubernetesNamespace
	}
	return ""
}

func (m *IncidentListRequest) GetKubernetesObjectType() string {
	if m != nil {
		return m.KubernetesObjectType
	}
	return ""
}

func (m *IncidentListRequest) GetKubernetesObjectName() string {
	if m != nil {
		return m.KubernetesObjectName
	}
	return ""
}

func (m *IncidentListRequest) GetStates() []string {
	if m != nil {
		return m.States
	}
	return nil
}

type IncidentListResponse struct {
	Incidents []*Incident `protobuf:"bytes,1,rep,name=incidents" json:"incidents,omitempty"`
}

func (m *IncidentListResponse) Reset()                    { *m = IncidentListResponse{} }
func (m *IncidentListResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentListResponse) ProtoMessage()               {}
func (*IncidentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *IncidentListResponse) GetIncidents() []*Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

// Next Id: 2
type IncidentDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *IncidentDescribeRequest) Reset()                    { *m = IncidentDescribeRequest{} }
func (m *IncidentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeRequest) ProtoMessage()               {}
func (*IncidentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *IncidentDescribeRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

type IncidentDescribeResponse struct {
	Incident *Incident `protobuf:"bytes,1,opt,name=incident" json:"incident,omitempty"`
}

func (m *IncidentDescribeResponse) Reset()                    { *m = IncidentDescribeResponse{} }
func (m *IncidentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeResponse) ProtoMessage()               {}
func (*IncidentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IncidentDescribeResponse) GetIncident() *Incident {
	if m != nil {
		return m.Incident
	}
	return nil
}

// Next Id: 12
type IncidentNotifyRequest struct {
	AlertPhid string `protobuf:"bytes,1,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State     string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Output    string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	// The time object is used in icinga to send request. This
	// indicates detection time from icinga.
	Time                int64  `protobuf:"varint,6,opt,name=time" json:"time,omitempty"`
	Author              string `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Comment             string `protobuf:"bytes,8,opt,name=comment" json:"comment,omitempty"`
	KubernetesAlertName string `protobuf:"bytes,9,opt,name=kubernetes_alert_name,json=kubernetesAlertName" json:"kubernetes_alert_name,omitempty"`
	KubernetesCluster   string `protobuf:"bytes,10,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
}

func (m *IncidentNotifyRequest) Reset()                    { *m = IncidentNotifyRequest{} }
func (m *IncidentNotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentNotifyRequest) ProtoMessage()               {}
func (*IncidentNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *IncidentNotifyRequest) GetAlertPhid() string {
	if m != nil {
		return m.AlertPhid
	}
	return ""
}

func (m *IncidentNotifyRequest) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *IncidentNotifyRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *IncidentNotifyRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *IncidentNotifyRequest) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *IncidentNotifyRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *IncidentNotifyRequest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *IncidentNotifyRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *IncidentNotifyRequest) GetKubernetesAlertName() string {
	if m != nil {
		return m.KubernetesAlertName
	}
	return ""
}

func (m *IncidentNotifyRequest) GetKubernetesCluster() string {
	if m != nil {
		return m.KubernetesCluster
	}
	return ""
}

// Next Id: 4
type IncidentEventCreateRequest struct {
	// Incident PHID
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Comment     string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Acknowledge bool   `protobuf:"varint,3,opt,name=acknowledge" json:"acknowledge,omitempty"`
}

func (m *IncidentEventCreateRequest) Reset()                    { *m = IncidentEventCreateRequest{} }
func (m *IncidentEventCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentEventCreateRequest) ProtoMessage()               {}
func (*IncidentEventCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *IncidentEventCreateRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *IncidentEventCreateRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *IncidentEventCreateRequest) GetAcknowledge() bool {
	if m != nil {
		return m.Acknowledge
	}
	return false
}

func init() {
	proto.RegisterType((*Incident)(nil), "appscode.kubernetes.v1beta1.Incident")
	proto.RegisterType((*Incident_Event)(nil), "appscode.kubernetes.v1beta1.Incident.Event")
	proto.RegisterType((*IncidentListRequest)(nil), "appscode.kubernetes.v1beta1.IncidentListRequest")
	proto.RegisterType((*IncidentListResponse)(nil), "appscode.kubernetes.v1beta1.IncidentListResponse")
	proto.RegisterType((*IncidentDescribeRequest)(nil), "appscode.kubernetes.v1beta1.IncidentDescribeRequest")
	proto.RegisterType((*IncidentDescribeResponse)(nil), "appscode.kubernetes.v1beta1.IncidentDescribeResponse")
	proto.RegisterType((*IncidentNotifyRequest)(nil), "appscode.kubernetes.v1beta1.IncidentNotifyRequest")
	proto.RegisterType((*IncidentEventCreateRequest)(nil), "appscode.kubernetes.v1beta1.IncidentEventCreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Incidents service

type IncidentsClient interface {
	List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error)
	Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error)
	Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type incidentsClient struct {
	cc *grpc.ClientConn
}

func NewIncidentsClient(cc *grpc.ClientConn) IncidentsClient {
	return &incidentsClient{cc}
}

func (c *incidentsClient) List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error) {
	out := new(IncidentListResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error) {
	out := new(IncidentDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.Incidents/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Incidents service

type IncidentsServer interface {
	List(context.Context, *IncidentListRequest) (*IncidentListResponse, error)
	Describe(context.Context, *IncidentDescribeRequest) (*IncidentDescribeResponse, error)
	Notify(context.Context, *IncidentNotifyRequest) (*appscode_dtypes.VoidResponse, error)
	CreateEvent(context.Context, *IncidentEventCreateRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterIncidentsServer(s *grpc.Server, srv IncidentsServer) {
	s.RegisterService(&_Incidents_serviceDesc, srv)
}

func _Incidents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).List(ctx, req.(*IncidentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Describe(ctx, req.(*IncidentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Notify(ctx, req.(*IncidentNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentEventCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.Incidents/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).CreateEvent(ctx, req.(*IncidentEventCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incidents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.kubernetes.v1beta1.Incidents",
	HandlerType: (*IncidentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Incidents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Incidents_Describe_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Incidents_Notify_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Incidents_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "incident.proto",
}

func init() { proto.RegisterFile("incident.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 944 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xd8, 0x8e, 0x63, 0x1f, 0x27, 0x29, 0x4c, 0xd3, 0x30, 0x72, 0x81, 0xba, 0x5b, 0x55,
	0x58, 0xa0, 0x78, 0x1b, 0x53, 0x54, 0xc1, 0x45, 0x91, 0x63, 0x2a, 0x81, 0x84, 0xda, 0x68, 0xa1,
	0x5c, 0xf0, 0x23, 0x6b, 0xbc, 0x3e, 0x24, 0xdb, 0xda, 0x3b, 0xcb, 0xce, 0x38, 0x55, 0x54, 0x55,
	0x42, 0x7d, 0x05, 0x24, 0x5e, 0x82, 0x6b, 0xb8, 0xe4, 0x22, 0x97, 0xdc, 0xf2, 0x0a, 0xbc, 0x05,
	0x37, 0xd5, 0xce, 0x8f, 0xbd, 0x8e, 0xed, 0xc4, 0xc9, 0x8d, 0xe5, 0x3d, 0x73, 0xbe, 0x99, 0xef,
	0x7c, 0xf3, 0x9d, 0xb3, 0x0b, 0x5b, 0x51, 0x1c, 0x46, 0x03, 0x8c, 0x55, 0x2b, 0x49, 0x85, 0x12,
	0xf4, 0x26, 0x4f, 0x12, 0x19, 0x8a, 0x01, 0xb6, 0x9e, 0x8f, 0xfb, 0x98, 0xc6, 0xa8, 0x50, 0xb6,
	0x8e, 0xf7, 0xfa, 0xa8, 0xf8, 0x5e, 0xfd, 0xdd, 0x43, 0x21, 0x0e, 0x87, 0xe8, 0xf3, 0x24, 0xf2,
	0x79, 0x1c, 0x0b, 0xc5, 0x55, 0x24, 0x62, 0x69, 0xa0, 0xf5, 0xf7, 0x1d, 0x74, 0xc9, 0xfa, 0xad,
	0x99, 0xf5, 0x81, 0x3a, 0x49, 0x50, 0xfa, 0xfa, 0xd7, 0x24, 0x78, 0xa7, 0x65, 0xa8, 0x7c, 0x65,
	0xe9, 0x50, 0x0a, 0xa5, 0xe4, 0x28, 0x1a, 0x30, 0xd2, 0x20, 0xcd, 0x6a, 0xa0, 0xff, 0xd3, 0x2e,
	0xd0, 0x29, 0xab, 0x5e, 0x38, 0x1c, 0x4b, 0x85, 0x29, 0x2b, 0x34, 0x48, 0xb3, 0xd6, 0xde, 0x6e,
	0x4d, 0x98, 0x9b, 0xad, 0x5b, 0x4f, 0xa3, 0x41, 0xf0, 0xf6, 0x34, 0xbf, 0x6b, 0xd2, 0xe9, 0x1e,
	0x6c, 0xe7, 0x36, 0x89, 0xf9, 0x08, 0x65, 0xc2, 0x43, 0x64, 0x45, 0x7d, 0xd0, 0xf5, 0xe9, 0xda,
	0x63, 0xb7, 0x44, 0xef, 0xc3, 0x4e, 0x0e, 0x22, 0xfa, 0xcf, 0x30, 0x54, 0xbd, 0xec, 0x10, 0x56,
	0xd2, 0xa0, 0xdc, 0x86, 0x4f, 0xf4, 0xe2, 0xb7, 0x27, 0xc9, 0x12, 0x54, 0x76, 0x1e, 0x5b, 0x5b,
	0x8c, 0xca, 0x0e, 0xa4, 0x6d, 0xb8, 0x91, 0x43, 0xf1, 0x21, 0xa6, 0x16, 0x54, 0x3e, 0xcb, 0xaf,
	0x93, 0xad, 0x69, 0xcc, 0x2d, 0xa8, 0x45, 0x61, 0x14, 0x1f, 0xf2, 0xde, 0x91, 0x90, 0x8a, 0xad,
	0xeb, 0x4c, 0x30, 0xa1, 0x2f, 0x85, 0x54, 0xf4, 0x2e, 0x6c, 0xd9, 0x04, 0x89, 0xe9, 0x71, 0x14,
	0x22, 0xab, 0xe8, 0x9c, 0x4d, 0x13, 0xfd, 0xc6, 0x04, 0x33, 0xcd, 0x75, 0x55, 0x55, 0xa3, 0x79,
	0xf6, 0x9f, 0x6e, 0xc3, 0x9a, 0x54, 0x5c, 0x21, 0x03, 0x1d, 0x34, 0x0f, 0xb4, 0x09, 0xa5, 0xb1,
	0xc4, 0x94, 0xd5, 0xce, 0xd1, 0x5e, 0x67, 0x64, 0xdc, 0x52, 0x4c, 0x44, 0xaa, 0x70, 0xd0, 0xe3,
	0x8a, 0x6d, 0x34, 0x48, 0xb3, 0x18, 0x80, 0x0b, 0x75, 0x14, 0xfd, 0x00, 0xae, 0xf1, 0xf0, 0x79,
	0x2c, 0x5e, 0x0c, 0x71, 0x70, 0x68, 0x92, 0x36, 0x75, 0xd2, 0x56, 0x3e, 0xdc, 0x51, 0xf4, 0x36,
	0x6c, 0xa4, 0x18, 0x8a, 0x63, 0x4c, 0x4d, 0xd6, 0x96, 0xce, 0xaa, 0x4d, 0x62, 0x1d, 0x45, 0xef,
	0x80, 0xad, 0xe8, 0x05, 0xf6, 0x7b, 0xe3, 0x74, 0xc8, 0xae, 0x69, 0xd2, 0x1b, 0x93, 0xe0, 0xd3,
	0x74, 0x48, 0xbb, 0x50, 0xc6, 0x63, 0x8c, 0x95, 0x64, 0x6f, 0x35, 0x8a, 0xcd, 0x5a, 0xfb, 0xa3,
	0xd6, 0x39, 0x9e, 0x6f, 0x39, 0x43, 0xb6, 0x1e, 0x65, 0x98, 0xc0, 0x42, 0xeb, 0xbf, 0x13, 0x58,
	0xd3, 0x91, 0x89, 0x68, 0x64, 0x91, 0x68, 0x85, 0xbc, 0x68, 0x67, 0xa4, 0x28, 0xce, 0x49, 0xe1,
	0x54, 0x2d, 0x5d, 0xa8, 0x2a, 0x83, 0xf5, 0x50, 0x8c, 0x46, 0x18, 0x2b, 0x6b, 0x26, 0xf7, 0xe8,
	0xfd, 0x4f, 0xe0, 0xba, 0xe3, 0xfc, 0x75, 0x24, 0x55, 0x80, 0xbf, 0x8c, 0x51, 0x2a, 0xba, 0xbb,
	0xb0, 0x77, 0x0c, 0xe9, 0x4b, 0x74, 0x49, 0xe1, 0x2a, 0x5d, 0x52, 0xbc, 0x52, 0x97, 0x94, 0xce,
	0xe9, 0x92, 0x1d, 0x28, 0x6b, 0x4d, 0x25, 0x5b, 0x6b, 0x14, 0x9b, 0xd5, 0xc0, 0x3e, 0x79, 0x3f,
	0xc0, 0xf6, 0x6c, 0xf1, 0x32, 0x11, 0xb1, 0x44, 0xda, 0x85, 0xaa, 0x1b, 0x74, 0x92, 0x11, 0x7d,
	0xed, 0x77, 0x57, 0xba, 0xf6, 0x60, 0x8a, 0xf3, 0x76, 0xe1, 0x1d, 0x17, 0xfe, 0x02, 0x65, 0x98,
	0x46, 0x7d, 0x74, 0xea, 0x2e, 0x98, 0x56, 0xde, 0x4f, 0xc0, 0xe6, 0xd3, 0x2d, 0x9f, 0x0e, 0x54,
	0xdc, 0xbe, 0x1a, 0xb3, 0x32, 0x9d, 0x09, 0xcc, 0x3b, 0x2d, 0xc0, 0x0d, 0x17, 0x7e, 0x2c, 0x54,
	0xf4, 0xf3, 0x89, 0x23, 0x73, 0x1b, 0xc0, 0xcc, 0x8d, 0x29, 0xa5, 0xfd, 0x02, 0x23, 0x41, 0x55,
	0x47, 0x0f, 0xb2, 0x49, 0x7a, 0x13, 0xaa, 0xd9, 0xa8, 0x30, 0x42, 0x9b, 0x3b, 0xad, 0x64, 0x01,
	0x2d, 0xae, 0x73, 0x74, 0x71, 0x91, 0xa3, 0x4b, 0x79, 0x47, 0xef, 0x40, 0x59, 0x8c, 0x55, 0x32,
	0x76, 0x2e, 0xb4, 0x4f, 0x7a, 0x87, 0xc8, 0xce, 0xac, 0x62, 0xa0, 0xff, 0x67, 0xb9, 0x7c, 0xac,
	0x8e, 0x44, 0x6a, 0xe7, 0x93, 0x7d, 0xca, 0x5b, 0xb9, 0x32, 0x63, 0xe5, 0xe5, 0xa3, 0xb0, 0xba,
	0x7c, 0x14, 0x2e, 0xb6, 0x39, 0x2c, 0xb1, 0xb9, 0x37, 0x84, 0xba, 0xd3, 0x50, 0x77, 0x73, 0x37,
	0x45, 0xae, 0xce, 0xbb, 0xd5, 0x3c, 0xdd, 0xc2, 0x2c, 0xdd, 0x06, 0xd4, 0x72, 0x13, 0x4b, 0xab,
	0x57, 0x09, 0xf2, 0xa1, 0xf6, 0xaf, 0x65, 0xa8, 0xba, 0xe3, 0x24, 0xfd, 0x83, 0x40, 0x29, 0x33,
	0x29, 0xbd, 0xb7, 0xd2, 0xd5, 0xe7, 0x9a, 0xb9, 0xbe, 0x77, 0x09, 0x84, 0x71, 0x9c, 0xf7, 0xf0,
	0xf5, 0x9f, 0xac, 0x50, 0x21, 0xaf, 0xff, 0xfd, 0xef, 0xb7, 0x42, 0x9b, 0xde, 0xf3, 0x7b, 0x33,
	0x6f, 0xe3, 0xe9, 0x2e, 0xbe, 0xdd, 0xc5, 0x9f, 0xf8, 0xde, 0x7f, 0x26, 0x45, 0x4c, 0xff, 0x26,
	0x50, 0x71, 0x36, 0xa6, 0xf7, 0x57, 0x3a, 0xff, 0x4c, 0x93, 0xd4, 0x3f, 0xb9, 0x24, 0xca, 0x32,
	0x7f, 0x94, 0x63, 0xfe, 0x29, 0x7d, 0x70, 0x19, 0xe6, 0x2f, 0xb3, 0x2b, 0x7b, 0x65, 0x0a, 0xf8,
	0x87, 0x40, 0xd9, 0xf4, 0x09, 0x6d, 0xaf, 0x44, 0x64, 0xa6, 0xa9, 0xea, 0xef, 0xcd, 0x4d, 0xe3,
	0xef, 0x44, 0x34, 0x98, 0x90, 0x1c, 0xe5, 0x48, 0x72, 0xef, 0xc7, 0x8b, 0x49, 0x5a, 0x6b, 0x4a,
	0xff, 0xe5, 0xbc, 0x5f, 0x5f, 0xf9, 0x3c, 0xd4, 0x5f, 0x4e, 0x7e, 0xac, 0x39, 0xec, 0xba, 0x8a,
	0x74, 0x25, 0x9f, 0x91, 0x0f, 0xe9, 0x5f, 0x04, 0x6a, 0xc6, 0xab, 0xe6, 0x25, 0xf4, 0x60, 0xa5,
	0x8a, 0xe6, 0x2d, 0x7e, 0x51, 0x59, 0x4f, 0x72, 0x65, 0x75, 0xbd, 0x87, 0x57, 0xd0, 0xde, 0xbc,
	0x2f, 0x1d, 0xf1, 0xfd, 0xcf, 0xe1, 0x4e, 0x28, 0x46, 0xd3, 0x43, 0x79, 0x12, 0x2d, 0x60, 0xbc,
	0xbf, 0xe9, 0x28, 0x1f, 0x64, 0x5f, 0x86, 0x07, 0xe4, 0xfb, 0x75, 0xbb, 0xd2, 0x2f, 0xeb, 0x6f,
	0xc5, 0x8f, 0xdf, 0x04, 0x00, 0x00, 0xff, 0xff, 0x54, 0x58, 0x98, 0x52, 0xb9, 0x0a, 0x00, 0x00,
}