// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloudapi/v1/cloudapi.proto

package cloudapiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInstancesRequest) Reset() {
	*x = ListInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstancesRequest) ProtoMessage() {}

func (x *ListInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstancesRequest.ProtoReflect.Descriptor instead.
func (*ListInstancesRequest) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{0}
}

type ListInstancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*Instance    `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	Drift     *InstanceDrift `protobuf:"bytes,2,opt,name=drift,proto3" json:"drift,omitempty"`
}

func (x *ListInstancesResponse) Reset() {
	*x = ListInstancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstancesResponse) ProtoMessage() {}

func (x *ListInstancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstancesResponse.ProtoReflect.Descriptor instead.
func (*ListInstancesResponse) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{1}
}

func (x *ListInstancesResponse) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *ListInstancesResponse) GetDrift() *InstanceDrift {
	if x != nil {
		return x.Drift
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are assignable to RepoDetails:
	//
	//	*Instance_V1RepoDetails
	//	*Instance_V2RepoDetails
	RepoDetails    isInstance_RepoDetails `protobuf_oneof:"repo_details"`
	ProjectDetails *ProjectDetails        `protobuf:"bytes,6,opt,name=project_details,json=projectDetails,proto3" json:"project_details,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{2}
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *Instance) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (m *Instance) GetRepoDetails() isInstance_RepoDetails {
	if m != nil {
		return m.RepoDetails
	}
	return nil
}

func (x *Instance) GetV1RepoDetails() *V1RepoDetails {
	if x, ok := x.GetRepoDetails().(*Instance_V1RepoDetails); ok {
		return x.V1RepoDetails
	}
	return nil
}

func (x *Instance) GetV2RepoDetails() *V2RepoDetails {
	if x, ok := x.GetRepoDetails().(*Instance_V2RepoDetails); ok {
		return x.V2RepoDetails
	}
	return nil
}

func (x *Instance) GetProjectDetails() *ProjectDetails {
	if x != nil {
		return x.ProjectDetails
	}
	return nil
}

type isInstance_RepoDetails interface {
	isInstance_RepoDetails()
}

type Instance_V1RepoDetails struct {
	V1RepoDetails *V1RepoDetails `protobuf:"bytes,4,opt,name=v1_repo_details,json=v1RepoDetails,proto3,oneof"`
}

type Instance_V2RepoDetails struct {
	V2RepoDetails *V2RepoDetails `protobuf:"bytes,5,opt,name=v2_repo_details,json=v2RepoDetails,proto3,oneof"`
}

func (*Instance_V1RepoDetails) isInstance_RepoDetails() {}

func (*Instance_V2RepoDetails) isInstance_RepoDetails() {}

type InstanceDrift struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlyRepo    []*Instance `protobuf:"bytes,1,rep,name=only_repo,json=onlyRepo,proto3" json:"only_repo,omitempty"`
	OnlyProject []*Instance `protobuf:"bytes,2,rep,name=only_project,json=onlyProject,proto3" json:"only_project,omitempty"`
}

func (x *InstanceDrift) Reset() {
	*x = InstanceDrift{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceDrift) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceDrift) ProtoMessage() {}

func (x *InstanceDrift) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceDrift.ProtoReflect.Descriptor instead.
func (*InstanceDrift) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceDrift) GetOnlyRepo() []*Instance {
	if x != nil {
		return x.OnlyRepo
	}
	return nil
}

func (x *InstanceDrift) GetOnlyProject() []*Instance {
	if x != nil {
		return x.OnlyProject
	}
	return nil
}

type ProjectDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId    string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	EmailDomain  string                 `protobuf:"bytes,2,opt,name=email_domain,json=emailDomain,proto3" json:"email_domain,omitempty"`
	InstanceType string                 `protobuf:"bytes,3,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	IsDeleted    bool                   `protobuf:"varint,4,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ProjectDetails) Reset() {
	*x = ProjectDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDetails) ProtoMessage() {}

func (x *ProjectDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDetails.ProtoReflect.Descriptor instead.
func (*ProjectDetails) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectDetails) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectDetails) GetEmailDomain() string {
	if x != nil {
		return x.EmailDomain
	}
	return ""
}

func (x *ProjectDetails) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *ProjectDetails) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ProjectDetails) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type V1RepoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer                           string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	EmailDomain                        string `protobuf:"bytes,2,opt,name=email_domain,json=emailDomain,proto3" json:"email_domain,omitempty"`
	Endpoint                           string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	InstanceType                       string `protobuf:"bytes,4,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Type                               string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DisableManagedSmtp                 bool   `protobuf:"varint,6,opt,name=disable_managed_smtp,json=disableManagedSmtp,proto3" json:"disable_managed_smtp,omitempty"`
	DisableSourcegraphManagementAccess bool   `protobuf:"varint,7,opt,name=disable_sourcegraph_management_access,json=disableSourcegraphManagementAccess,proto3" json:"disable_sourcegraph_management_access,omitempty"`
}

func (x *V1RepoDetails) Reset() {
	*x = V1RepoDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1RepoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1RepoDetails) ProtoMessage() {}

func (x *V1RepoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1RepoDetails.ProtoReflect.Descriptor instead.
func (*V1RepoDetails) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{5}
}

func (x *V1RepoDetails) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *V1RepoDetails) GetEmailDomain() string {
	if x != nil {
		return x.EmailDomain
	}
	return ""
}

func (x *V1RepoDetails) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *V1RepoDetails) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *V1RepoDetails) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *V1RepoDetails) GetDisableManagedSmtp() bool {
	if x != nil {
		return x.DisableManagedSmtp
	}
	return false
}

func (x *V1RepoDetails) GetDisableSourcegraphManagementAccess() bool {
	if x != nil {
		return x.DisableSourcegraphManagementAccess
	}
	return false
}

type V2RepoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *V2RepoDetails) Reset() {
	*x = V2RepoDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V2RepoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V2RepoDetails) ProtoMessage() {}

func (x *V2RepoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cloudapi_v1_cloudapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V2RepoDetails.ProtoReflect.Descriptor instead.
func (*V2RepoDetails) Descriptor() ([]byte, []int) {
	return file_cloudapi_v1_cloudapi_proto_rawDescGZIP(), []int{6}
}

var File_cloudapi_v1_cloudapi_proto protoreflect.FileDescriptor

var file_cloudapi_v1_cloudapi_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x7e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x30, 0x0a, 0x05, 0x64, 0x72, 0x69, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x72, 0x69, 0x66, 0x74, 0x52, 0x05, 0x64, 0x72, 0x69,
	0x66, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x76, 0x31, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x31,
	0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x76,
	0x31, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0f,
	0x76, 0x32, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x32, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x48, 0x00, 0x52, 0x0d, 0x76, 0x32, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x7d, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x44, 0x72, 0x69, 0x66, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x6f, 0x6e, 0x6c,
	0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x38, 0x0a,
	0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x79,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x02, 0x0a, 0x0d,
	0x56, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x73, 0x6d, 0x74, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53,
	0x6d, 0x74, 0x70, 0x12, 0x51, 0x0a, 0x25, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x22, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x56, 0x32, 0x52, 0x65, 0x70, 0x6f,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x6b, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudapi_v1_cloudapi_proto_rawDescOnce sync.Once
	file_cloudapi_v1_cloudapi_proto_rawDescData = file_cloudapi_v1_cloudapi_proto_rawDesc
)

func file_cloudapi_v1_cloudapi_proto_rawDescGZIP() []byte {
	file_cloudapi_v1_cloudapi_proto_rawDescOnce.Do(func() {
		file_cloudapi_v1_cloudapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudapi_v1_cloudapi_proto_rawDescData)
	})
	return file_cloudapi_v1_cloudapi_proto_rawDescData
}

var file_cloudapi_v1_cloudapi_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloudapi_v1_cloudapi_proto_goTypes = []interface{}{
	(*ListInstancesRequest)(nil),  // 0: cloudapi.v1.ListInstancesRequest
	(*ListInstancesResponse)(nil), // 1: cloudapi.v1.ListInstancesResponse
	(*Instance)(nil),              // 2: cloudapi.v1.Instance
	(*InstanceDrift)(nil),         // 3: cloudapi.v1.InstanceDrift
	(*ProjectDetails)(nil),        // 4: cloudapi.v1.ProjectDetails
	(*V1RepoDetails)(nil),         // 5: cloudapi.v1.V1RepoDetails
	(*V2RepoDetails)(nil),         // 6: cloudapi.v1.V2RepoDetails
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_cloudapi_v1_cloudapi_proto_depIdxs = []int32{
	2, // 0: cloudapi.v1.ListInstancesResponse.instances:type_name -> cloudapi.v1.Instance
	3, // 1: cloudapi.v1.ListInstancesResponse.drift:type_name -> cloudapi.v1.InstanceDrift
	5, // 2: cloudapi.v1.Instance.v1_repo_details:type_name -> cloudapi.v1.V1RepoDetails
	6, // 3: cloudapi.v1.Instance.v2_repo_details:type_name -> cloudapi.v1.V2RepoDetails
	4, // 4: cloudapi.v1.Instance.project_details:type_name -> cloudapi.v1.ProjectDetails
	2, // 5: cloudapi.v1.InstanceDrift.only_repo:type_name -> cloudapi.v1.Instance
	2, // 6: cloudapi.v1.InstanceDrift.only_project:type_name -> cloudapi.v1.Instance
	7, // 7: cloudapi.v1.ProjectDetails.created_at:type_name -> google.protobuf.Timestamp
	0, // 8: cloudapi.v1.InstanceService.ListInstances:input_type -> cloudapi.v1.ListInstancesRequest
	1, // 9: cloudapi.v1.InstanceService.ListInstances:output_type -> cloudapi.v1.ListInstancesResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloudapi_v1_cloudapi_proto_init() }
func file_cloudapi_v1_cloudapi_proto_init() {
	if File_cloudapi_v1_cloudapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudapi_v1_cloudapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstancesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstancesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceDrift); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1RepoDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloudapi_v1_cloudapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V2RepoDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cloudapi_v1_cloudapi_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Instance_V1RepoDetails)(nil),
		(*Instance_V2RepoDetails)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudapi_v1_cloudapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudapi_v1_cloudapi_proto_goTypes,
		DependencyIndexes: file_cloudapi_v1_cloudapi_proto_depIdxs,
		MessageInfos:      file_cloudapi_v1_cloudapi_proto_msgTypes,
	}.Build()
	File_cloudapi_v1_cloudapi_proto = out.File
	file_cloudapi_v1_cloudapi_proto_rawDesc = nil
	file_cloudapi_v1_cloudapi_proto_goTypes = nil
	file_cloudapi_v1_cloudapi_proto_depIdxs = nil
}
