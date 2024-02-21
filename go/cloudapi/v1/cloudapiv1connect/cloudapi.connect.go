// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudapi/v1/cloudapi.proto

package cloudapiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sourcegraph/cloud-api/go/cloudapi/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// InstanceServiceName is the fully-qualified name of the InstanceService service.
	InstanceServiceName = "cloudapi.v1.InstanceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InstanceServiceListInstancesProcedure is the fully-qualified name of the InstanceService's
	// ListInstances RPC.
	InstanceServiceListInstancesProcedure = "/cloudapi.v1.InstanceService/ListInstances"
	// InstanceServiceCreateOrUpdateInstanceProcedure is the fully-qualified name of the
	// InstanceService's CreateOrUpdateInstance RPC.
	InstanceServiceCreateOrUpdateInstanceProcedure = "/cloudapi.v1.InstanceService/CreateOrUpdateInstance"
	// InstanceServiceGetInstanceProcedure is the fully-qualified name of the InstanceService's
	// GetInstance RPC.
	InstanceServiceGetInstanceProcedure = "/cloudapi.v1.InstanceService/GetInstance"
)

// InstanceServiceClient is a client for the cloudapi.v1.InstanceService service.
type InstanceServiceClient interface {
	ListInstances(context.Context, *connect.Request[v1.ListInstancesRequest]) (*connect.Response[v1.ListInstancesResponse], error)
	CreateOrUpdateInstance(context.Context, *connect.Request[v1.CreateOrUpdateInstanceRequest]) (*connect.Response[v1.CreateOrUpdateInstanceResponse], error)
	GetInstance(context.Context, *connect.Request[v1.GetInstanceRequest]) (*connect.Response[v1.GetInstanceResponse], error)
}

// NewInstanceServiceClient constructs a client for the cloudapi.v1.InstanceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInstanceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) InstanceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &instanceServiceClient{
		listInstances: connect.NewClient[v1.ListInstancesRequest, v1.ListInstancesResponse](
			httpClient,
			baseURL+InstanceServiceListInstancesProcedure,
			opts...,
		),
		createOrUpdateInstance: connect.NewClient[v1.CreateOrUpdateInstanceRequest, v1.CreateOrUpdateInstanceResponse](
			httpClient,
			baseURL+InstanceServiceCreateOrUpdateInstanceProcedure,
			opts...,
		),
		getInstance: connect.NewClient[v1.GetInstanceRequest, v1.GetInstanceResponse](
			httpClient,
			baseURL+InstanceServiceGetInstanceProcedure,
			opts...,
		),
	}
}

// instanceServiceClient implements InstanceServiceClient.
type instanceServiceClient struct {
	listInstances          *connect.Client[v1.ListInstancesRequest, v1.ListInstancesResponse]
	createOrUpdateInstance *connect.Client[v1.CreateOrUpdateInstanceRequest, v1.CreateOrUpdateInstanceResponse]
	getInstance            *connect.Client[v1.GetInstanceRequest, v1.GetInstanceResponse]
}

// ListInstances calls cloudapi.v1.InstanceService.ListInstances.
func (c *instanceServiceClient) ListInstances(ctx context.Context, req *connect.Request[v1.ListInstancesRequest]) (*connect.Response[v1.ListInstancesResponse], error) {
	return c.listInstances.CallUnary(ctx, req)
}

// CreateOrUpdateInstance calls cloudapi.v1.InstanceService.CreateOrUpdateInstance.
func (c *instanceServiceClient) CreateOrUpdateInstance(ctx context.Context, req *connect.Request[v1.CreateOrUpdateInstanceRequest]) (*connect.Response[v1.CreateOrUpdateInstanceResponse], error) {
	return c.createOrUpdateInstance.CallUnary(ctx, req)
}

// GetInstance calls cloudapi.v1.InstanceService.GetInstance.
func (c *instanceServiceClient) GetInstance(ctx context.Context, req *connect.Request[v1.GetInstanceRequest]) (*connect.Response[v1.GetInstanceResponse], error) {
	return c.getInstance.CallUnary(ctx, req)
}

// InstanceServiceHandler is an implementation of the cloudapi.v1.InstanceService service.
type InstanceServiceHandler interface {
	ListInstances(context.Context, *connect.Request[v1.ListInstancesRequest]) (*connect.Response[v1.ListInstancesResponse], error)
	CreateOrUpdateInstance(context.Context, *connect.Request[v1.CreateOrUpdateInstanceRequest]) (*connect.Response[v1.CreateOrUpdateInstanceResponse], error)
	GetInstance(context.Context, *connect.Request[v1.GetInstanceRequest]) (*connect.Response[v1.GetInstanceResponse], error)
}

// NewInstanceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInstanceServiceHandler(svc InstanceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	instanceServiceListInstancesHandler := connect.NewUnaryHandler(
		InstanceServiceListInstancesProcedure,
		svc.ListInstances,
		opts...,
	)
	instanceServiceCreateOrUpdateInstanceHandler := connect.NewUnaryHandler(
		InstanceServiceCreateOrUpdateInstanceProcedure,
		svc.CreateOrUpdateInstance,
		opts...,
	)
	instanceServiceGetInstanceHandler := connect.NewUnaryHandler(
		InstanceServiceGetInstanceProcedure,
		svc.GetInstance,
		opts...,
	)
	return "/cloudapi.v1.InstanceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InstanceServiceListInstancesProcedure:
			instanceServiceListInstancesHandler.ServeHTTP(w, r)
		case InstanceServiceCreateOrUpdateInstanceProcedure:
			instanceServiceCreateOrUpdateInstanceHandler.ServeHTTP(w, r)
		case InstanceServiceGetInstanceProcedure:
			instanceServiceGetInstanceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInstanceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedInstanceServiceHandler struct{}

func (UnimplementedInstanceServiceHandler) ListInstances(context.Context, *connect.Request[v1.ListInstancesRequest]) (*connect.Response[v1.ListInstancesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudapi.v1.InstanceService.ListInstances is not implemented"))
}

func (UnimplementedInstanceServiceHandler) CreateOrUpdateInstance(context.Context, *connect.Request[v1.CreateOrUpdateInstanceRequest]) (*connect.Response[v1.CreateOrUpdateInstanceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudapi.v1.InstanceService.CreateOrUpdateInstance is not implemented"))
}

func (UnimplementedInstanceServiceHandler) GetInstance(context.Context, *connect.Request[v1.GetInstanceRequest]) (*connect.Response[v1.GetInstanceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudapi.v1.InstanceService.GetInstance is not implemented"))
}
