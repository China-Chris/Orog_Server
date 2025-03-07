// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: order.proto

package orderclient

import (
	"context"

	"github.com/simance-ai/smdx/rpcx/order/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderRequest       = order.CreateOrderRequest
	CreateOrderResponse      = order.CreateOrderResponse
	GetAccountInfoRequest    = order.GetAccountInfoRequest
	GetAccountInfoResponse   = order.GetAccountInfoResponse
	GetOrderDetailRequest    = order.GetOrderDetailRequest
	GetOrderDetailResponse   = order.GetOrderDetailResponse
	GetOrderListRequest      = order.GetOrderListRequest
	OrderListResponse        = order.OrderListResponse
	Orders                   = order.Orders
	QueryOrderRequest        = order.QueryOrderRequest
	QueryOrderResponse       = order.QueryOrderResponse
	RebateOrder              = order.RebateOrder
	RebateOrderListRequest   = order.RebateOrderListRequest
	RebateOrderListResponse  = order.RebateOrderListResponse
	Request                  = order.Request
	Response                 = order.Response
	UpdateRebateOrderRequest = order.UpdateRebateOrderRequest

	Order interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error)
		GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error)
		GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
		GetRebateOrderList(ctx context.Context, in *RebateOrderListRequest, opts ...grpc.CallOption) (*RebateOrderListResponse, error)
		UpdateRebateOrder(ctx context.Context, in *UpdateRebateOrderRequest, opts ...grpc.CallOption) (*Response, error)
		GetAccountOrderInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) QueryOrder(ctx context.Context, in *QueryOrderRequest, opts ...grpc.CallOption) (*QueryOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.QueryOrder(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetOrderDetail(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetOrderList(ctx, in, opts...)
}

func (m *defaultOrder) GetRebateOrderList(ctx context.Context, in *RebateOrderListRequest, opts ...grpc.CallOption) (*RebateOrderListResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetRebateOrderList(ctx, in, opts...)
}

func (m *defaultOrder) UpdateRebateOrder(ctx context.Context, in *UpdateRebateOrderRequest, opts ...grpc.CallOption) (*Response, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.UpdateRebateOrder(ctx, in, opts...)
}

func (m *defaultOrder) GetAccountOrderInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetAccountOrderInfo(ctx, in, opts...)
}
