// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wtclientrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WatchtowerClientClient is the client API for WatchtowerClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchtowerClientClient interface {
	// lncli: `wtclient add`
	// AddTower adds a new watchtower reachable at the given address and
	// considers it for new sessions. If the watchtower already exists, then
	// any new addresses included will be considered when dialing it for
	// session negotiations and backups.
	AddTower(ctx context.Context, in *AddTowerRequest, opts ...grpc.CallOption) (*AddTowerResponse, error)
	// lncli: `wtclient remove`
	// RemoveTower removes a watchtower from being considered for future session
	// negotiations and from being used for any subsequent backups until it's added
	// again. If an address is provided, then this RPC only serves as a way of
	// removing the address from the watchtower instead.
	RemoveTower(ctx context.Context, in *RemoveTowerRequest, opts ...grpc.CallOption) (*RemoveTowerResponse, error)
	// lncli: `wtclient deactivate`
	// DeactivateTower sets the given tower's status to inactive so that it
	// is not considered for session negotiation. Its sessions will also not
	// be used while the tower is inactive.
	DeactivateTower(ctx context.Context, in *DeactivateTowerRequest, opts ...grpc.CallOption) (*DeactivateTowerResponse, error)
	// lncli: `wtclient towers`
	// ListTowers returns the list of watchtowers registered with the client.
	ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (*ListTowersResponse, error)
	// lncli: `wtclient tower`
	// GetTowerInfo retrieves information for a registered watchtower.
	GetTowerInfo(ctx context.Context, in *GetTowerInfoRequest, opts ...grpc.CallOption) (*Tower, error)
	// lncli: `wtclient stats`
	// Stats returns the in-memory statistics of the client since startup.
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	// lncli: `wtclient policy`
	// Policy returns the active watchtower client policy configuration.
	Policy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyResponse, error)
}

type watchtowerClientClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchtowerClientClient(cc grpc.ClientConnInterface) WatchtowerClientClient {
	return &watchtowerClientClient{cc}
}

func (c *watchtowerClientClient) AddTower(ctx context.Context, in *AddTowerRequest, opts ...grpc.CallOption) (*AddTowerResponse, error) {
	out := new(AddTowerResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/AddTower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) RemoveTower(ctx context.Context, in *RemoveTowerRequest, opts ...grpc.CallOption) (*RemoveTowerResponse, error) {
	out := new(RemoveTowerResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/RemoveTower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) DeactivateTower(ctx context.Context, in *DeactivateTowerRequest, opts ...grpc.CallOption) (*DeactivateTowerResponse, error) {
	out := new(DeactivateTowerResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/DeactivateTower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (*ListTowersResponse, error) {
	out := new(ListTowersResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/ListTowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) GetTowerInfo(ctx context.Context, in *GetTowerInfoRequest, opts ...grpc.CallOption) (*Tower, error) {
	out := new(Tower)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/GetTowerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchtowerClientClient) Policy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyResponse, error) {
	out := new(PolicyResponse)
	err := c.cc.Invoke(ctx, "/wtclientrpc.WatchtowerClient/Policy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchtowerClientServer is the server API for WatchtowerClient service.
// All implementations must embed UnimplementedWatchtowerClientServer
// for forward compatibility
type WatchtowerClientServer interface {
	// lncli: `wtclient add`
	// AddTower adds a new watchtower reachable at the given address and
	// considers it for new sessions. If the watchtower already exists, then
	// any new addresses included will be considered when dialing it for
	// session negotiations and backups.
	AddTower(context.Context, *AddTowerRequest) (*AddTowerResponse, error)
	// lncli: `wtclient remove`
	// RemoveTower removes a watchtower from being considered for future session
	// negotiations and from being used for any subsequent backups until it's added
	// again. If an address is provided, then this RPC only serves as a way of
	// removing the address from the watchtower instead.
	RemoveTower(context.Context, *RemoveTowerRequest) (*RemoveTowerResponse, error)
	// lncli: `wtclient deactivate`
	// DeactivateTower sets the given tower's status to inactive so that it
	// is not considered for session negotiation. Its sessions will also not
	// be used while the tower is inactive.
	DeactivateTower(context.Context, *DeactivateTowerRequest) (*DeactivateTowerResponse, error)
	// lncli: `wtclient towers`
	// ListTowers returns the list of watchtowers registered with the client.
	ListTowers(context.Context, *ListTowersRequest) (*ListTowersResponse, error)
	// lncli: `wtclient tower`
	// GetTowerInfo retrieves information for a registered watchtower.
	GetTowerInfo(context.Context, *GetTowerInfoRequest) (*Tower, error)
	// lncli: `wtclient stats`
	// Stats returns the in-memory statistics of the client since startup.
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	// lncli: `wtclient policy`
	// Policy returns the active watchtower client policy configuration.
	Policy(context.Context, *PolicyRequest) (*PolicyResponse, error)
	mustEmbedUnimplementedWatchtowerClientServer()
}

// UnimplementedWatchtowerClientServer must be embedded to have forward compatible implementations.
type UnimplementedWatchtowerClientServer struct {
}

func (UnimplementedWatchtowerClientServer) AddTower(context.Context, *AddTowerRequest) (*AddTowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTower not implemented")
}
func (UnimplementedWatchtowerClientServer) RemoveTower(context.Context, *RemoveTowerRequest) (*RemoveTowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTower not implemented")
}
func (UnimplementedWatchtowerClientServer) DeactivateTower(context.Context, *DeactivateTowerRequest) (*DeactivateTowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateTower not implemented")
}
func (UnimplementedWatchtowerClientServer) ListTowers(context.Context, *ListTowersRequest) (*ListTowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTowers not implemented")
}
func (UnimplementedWatchtowerClientServer) GetTowerInfo(context.Context, *GetTowerInfoRequest) (*Tower, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTowerInfo not implemented")
}
func (UnimplementedWatchtowerClientServer) Stats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedWatchtowerClientServer) Policy(context.Context, *PolicyRequest) (*PolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Policy not implemented")
}
func (UnimplementedWatchtowerClientServer) mustEmbedUnimplementedWatchtowerClientServer() {}

// UnsafeWatchtowerClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchtowerClientServer will
// result in compilation errors.
type UnsafeWatchtowerClientServer interface {
	mustEmbedUnimplementedWatchtowerClientServer()
}

func RegisterWatchtowerClientServer(s grpc.ServiceRegistrar, srv WatchtowerClientServer) {
	s.RegisterService(&WatchtowerClient_ServiceDesc, srv)
}

func _WatchtowerClient_AddTower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).AddTower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/AddTower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).AddTower(ctx, req.(*AddTowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_RemoveTower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).RemoveTower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/RemoveTower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).RemoveTower(ctx, req.(*RemoveTowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_DeactivateTower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateTowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).DeactivateTower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/DeactivateTower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).DeactivateTower(ctx, req.(*DeactivateTowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_ListTowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).ListTowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/ListTowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).ListTowers(ctx, req.(*ListTowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_GetTowerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTowerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).GetTowerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/GetTowerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).GetTowerInfo(ctx, req.(*GetTowerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WatchtowerClient_Policy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerClientServer).Policy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wtclientrpc.WatchtowerClient/Policy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerClientServer).Policy(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WatchtowerClient_ServiceDesc is the grpc.ServiceDesc for WatchtowerClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchtowerClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wtclientrpc.WatchtowerClient",
	HandlerType: (*WatchtowerClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTower",
			Handler:    _WatchtowerClient_AddTower_Handler,
		},
		{
			MethodName: "RemoveTower",
			Handler:    _WatchtowerClient_RemoveTower_Handler,
		},
		{
			MethodName: "DeactivateTower",
			Handler:    _WatchtowerClient_DeactivateTower_Handler,
		},
		{
			MethodName: "ListTowers",
			Handler:    _WatchtowerClient_ListTowers_Handler,
		},
		{
			MethodName: "GetTowerInfo",
			Handler:    _WatchtowerClient_GetTowerInfo_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _WatchtowerClient_Stats_Handler,
		},
		{
			MethodName: "Policy",
			Handler:    _WatchtowerClient_Policy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wtclientrpc/wtclient.proto",
}
