// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: chains/sol.proto

package sol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Sol_Ping_FullMethodName                     = "/sol.Sol/Ping"
	Sol_GetTxByHash_FullMethodName              = "/sol.Sol/GetTxByHash"
	Sol_Start_FullMethodName                    = "/sol.Sol/Start"
	Sol_Trade_FullMethodName                    = "/sol.Sol/Trade"
	Sol_GetMarketInfo_FullMethodName            = "/sol.Sol/GetMarketInfo"
	Sol_GetMarketInfoByQuoteMint_FullMethodName = "/sol.Sol/GetMarketInfoByQuoteMint"
	Sol_GetMarketList_FullMethodName            = "/sol.Sol/GetMarketList"
	Sol_GetMarketKlineList_FullMethodName       = "/sol.Sol/GetMarketKlineList"
	Sol_GetTrendingMarket_FullMethodName        = "/sol.Sol/GetTrendingMarket"
	Sol_GetXTopMarketList_FullMethodName        = "/sol.Sol/GetXTopMarketList"
	Sol_GetMemeMarketList_FullMethodName        = "/sol.Sol/GetMemeMarketList"
	Sol_GetFollowMarketList_FullMethodName      = "/sol.Sol/GetFollowMarketList"
	Sol_SearchCurrencyList_FullMethodName       = "/sol.Sol/SearchCurrencyList"
	Sol_MarketActivityList_FullMethodName       = "/sol.Sol/MarketActivityList"
	Sol_GetMarketDetail_FullMethodName          = "/sol.Sol/GetMarketDetail"
)

// SolClient is the client API for Sol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetTxByHash(ctx context.Context, in *GetTxByHashRequest, opts ...grpc.CallOption) (*Tx, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*Response, error)
	Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*Response, error)
	GetMarketInfo(ctx context.Context, in *GetMarketInfoRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error)
	GetMarketInfoByQuoteMint(ctx context.Context, in *GetMarketInfoByQuoteMintRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error)
	GetMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketListResponse, error)
	GetMarketKlineList(ctx context.Context, in *GetMarketKlineRequest, opts ...grpc.CallOption) (*MarketKlineListResponse, error)
	GetTrendingMarket(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error)
	GetXTopMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error)
	GetMemeMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error)
	GetFollowMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error)
	SearchCurrencyList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*SearchCurrencyListResponse, error)
	MarketActivityList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketActivityListResponse, error)
	GetMarketDetail(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketDetailResponse, error)
}

type solClient struct {
	cc grpc.ClientConnInterface
}

func NewSolClient(cc grpc.ClientConnInterface) SolClient {
	return &solClient{cc}
}

func (c *solClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Sol_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetTxByHash(ctx context.Context, in *GetTxByHashRequest, opts ...grpc.CallOption) (*Tx, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tx)
	err := c.cc.Invoke(ctx, Sol_GetTxByHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Sol_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Sol_Trade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMarketInfo(ctx context.Context, in *GetMarketInfoRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketInfoResponse)
	err := c.cc.Invoke(ctx, Sol_GetMarketInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMarketInfoByQuoteMint(ctx context.Context, in *GetMarketInfoByQuoteMintRequest, opts ...grpc.CallOption) (*MarketInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketInfoResponse)
	err := c.cc.Invoke(ctx, Sol_GetMarketInfoByQuoteMint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketListResponse)
	err := c.cc.Invoke(ctx, Sol_GetMarketList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMarketKlineList(ctx context.Context, in *GetMarketKlineRequest, opts ...grpc.CallOption) (*MarketKlineListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketKlineListResponse)
	err := c.cc.Invoke(ctx, Sol_GetMarketKlineList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetTrendingMarket(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealTimeMarketListResponse)
	err := c.cc.Invoke(ctx, Sol_GetTrendingMarket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetXTopMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealTimeMarketListResponse)
	err := c.cc.Invoke(ctx, Sol_GetXTopMarketList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMemeMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealTimeMarketListResponse)
	err := c.cc.Invoke(ctx, Sol_GetMemeMarketList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetFollowMarketList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*RealTimeMarketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RealTimeMarketListResponse)
	err := c.cc.Invoke(ctx, Sol_GetFollowMarketList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) SearchCurrencyList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*SearchCurrencyListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCurrencyListResponse)
	err := c.cc.Invoke(ctx, Sol_SearchCurrencyList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) MarketActivityList(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketActivityListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketActivityListResponse)
	err := c.cc.Invoke(ctx, Sol_MarketActivityList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solClient) GetMarketDetail(ctx context.Context, in *GetMarketListRequest, opts ...grpc.CallOption) (*MarketDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketDetailResponse)
	err := c.cc.Invoke(ctx, Sol_GetMarketDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolServer is the server API for Sol service.
// All implementations must embed UnimplementedSolServer
// for forward compatibility.
type SolServer interface {
	Ping(context.Context, *Request) (*Response, error)
	GetTxByHash(context.Context, *GetTxByHashRequest) (*Tx, error)
	Start(context.Context, *StartRequest) (*Response, error)
	Trade(context.Context, *TradeRequest) (*Response, error)
	GetMarketInfo(context.Context, *GetMarketInfoRequest) (*MarketInfoResponse, error)
	GetMarketInfoByQuoteMint(context.Context, *GetMarketInfoByQuoteMintRequest) (*MarketInfoResponse, error)
	GetMarketList(context.Context, *GetMarketListRequest) (*MarketListResponse, error)
	GetMarketKlineList(context.Context, *GetMarketKlineRequest) (*MarketKlineListResponse, error)
	GetTrendingMarket(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error)
	GetXTopMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error)
	GetMemeMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error)
	GetFollowMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error)
	SearchCurrencyList(context.Context, *GetMarketListRequest) (*SearchCurrencyListResponse, error)
	MarketActivityList(context.Context, *GetMarketListRequest) (*MarketActivityListResponse, error)
	GetMarketDetail(context.Context, *GetMarketListRequest) (*MarketDetailResponse, error)
	mustEmbedUnimplementedSolServer()
}

// UnimplementedSolServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSolServer struct{}

func (UnimplementedSolServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSolServer) GetTxByHash(context.Context, *GetTxByHashRequest) (*Tx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByHash not implemented")
}
func (UnimplementedSolServer) Start(context.Context, *StartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSolServer) Trade(context.Context, *TradeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trade not implemented")
}
func (UnimplementedSolServer) GetMarketInfo(context.Context, *GetMarketInfoRequest) (*MarketInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketInfo not implemented")
}
func (UnimplementedSolServer) GetMarketInfoByQuoteMint(context.Context, *GetMarketInfoByQuoteMintRequest) (*MarketInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketInfoByQuoteMint not implemented")
}
func (UnimplementedSolServer) GetMarketList(context.Context, *GetMarketListRequest) (*MarketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketList not implemented")
}
func (UnimplementedSolServer) GetMarketKlineList(context.Context, *GetMarketKlineRequest) (*MarketKlineListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketKlineList not implemented")
}
func (UnimplementedSolServer) GetTrendingMarket(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrendingMarket not implemented")
}
func (UnimplementedSolServer) GetXTopMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetXTopMarketList not implemented")
}
func (UnimplementedSolServer) GetMemeMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemeMarketList not implemented")
}
func (UnimplementedSolServer) GetFollowMarketList(context.Context, *GetMarketListRequest) (*RealTimeMarketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowMarketList not implemented")
}
func (UnimplementedSolServer) SearchCurrencyList(context.Context, *GetMarketListRequest) (*SearchCurrencyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCurrencyList not implemented")
}
func (UnimplementedSolServer) MarketActivityList(context.Context, *GetMarketListRequest) (*MarketActivityListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketActivityList not implemented")
}
func (UnimplementedSolServer) GetMarketDetail(context.Context, *GetMarketListRequest) (*MarketDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketDetail not implemented")
}
func (UnimplementedSolServer) mustEmbedUnimplementedSolServer() {}
func (UnimplementedSolServer) testEmbeddedByValue()             {}

// UnsafeSolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolServer will
// result in compilation errors.
type UnsafeSolServer interface {
	mustEmbedUnimplementedSolServer()
}

func RegisterSolServer(s grpc.ServiceRegistrar, srv SolServer) {
	// If the following call pancis, it indicates UnimplementedSolServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sol_ServiceDesc, srv)
}

func _Sol_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetTxByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetTxByHash(ctx, req.(*GetTxByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_Trade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).Trade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_Trade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).Trade(ctx, req.(*TradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMarketInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMarketInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMarketInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMarketInfo(ctx, req.(*GetMarketInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMarketInfoByQuoteMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketInfoByQuoteMintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMarketInfoByQuoteMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMarketInfoByQuoteMint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMarketInfoByQuoteMint(ctx, req.(*GetMarketInfoByQuoteMintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMarketList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMarketList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMarketList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMarketList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMarketKlineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketKlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMarketKlineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMarketKlineList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMarketKlineList(ctx, req.(*GetMarketKlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetTrendingMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetTrendingMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetTrendingMarket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetTrendingMarket(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetXTopMarketList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetXTopMarketList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetXTopMarketList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetXTopMarketList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMemeMarketList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMemeMarketList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMemeMarketList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMemeMarketList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetFollowMarketList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetFollowMarketList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetFollowMarketList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetFollowMarketList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_SearchCurrencyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).SearchCurrencyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_SearchCurrencyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).SearchCurrencyList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_MarketActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).MarketActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_MarketActivityList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).MarketActivityList(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sol_GetMarketDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolServer).GetMarketDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sol_GetMarketDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolServer).GetMarketDetail(ctx, req.(*GetMarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sol_ServiceDesc is the grpc.ServiceDesc for Sol service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sol_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sol.Sol",
	HandlerType: (*SolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Sol_Ping_Handler,
		},
		{
			MethodName: "GetTxByHash",
			Handler:    _Sol_GetTxByHash_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Sol_Start_Handler,
		},
		{
			MethodName: "Trade",
			Handler:    _Sol_Trade_Handler,
		},
		{
			MethodName: "GetMarketInfo",
			Handler:    _Sol_GetMarketInfo_Handler,
		},
		{
			MethodName: "GetMarketInfoByQuoteMint",
			Handler:    _Sol_GetMarketInfoByQuoteMint_Handler,
		},
		{
			MethodName: "GetMarketList",
			Handler:    _Sol_GetMarketList_Handler,
		},
		{
			MethodName: "GetMarketKlineList",
			Handler:    _Sol_GetMarketKlineList_Handler,
		},
		{
			MethodName: "GetTrendingMarket",
			Handler:    _Sol_GetTrendingMarket_Handler,
		},
		{
			MethodName: "GetXTopMarketList",
			Handler:    _Sol_GetXTopMarketList_Handler,
		},
		{
			MethodName: "GetMemeMarketList",
			Handler:    _Sol_GetMemeMarketList_Handler,
		},
		{
			MethodName: "GetFollowMarketList",
			Handler:    _Sol_GetFollowMarketList_Handler,
		},
		{
			MethodName: "SearchCurrencyList",
			Handler:    _Sol_SearchCurrencyList_Handler,
		},
		{
			MethodName: "MarketActivityList",
			Handler:    _Sol_MarketActivityList_Handler,
		},
		{
			MethodName: "GetMarketDetail",
			Handler:    _Sol_GetMarketDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chains/sol.proto",
}
