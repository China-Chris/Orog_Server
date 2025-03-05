// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: sol.proto

package server

import (
	"context"

	"github.com/simance-ai/smdx/rpcx/chains/sol/internal/logic"
	"github.com/simance-ai/smdx/rpcx/chains/sol/internal/svc"
	"github.com/simance-ai/smdx/rpcx/chains/sol/sol"
)

type SolServer struct {
	svcCtx *svc.ServiceContext
	sol.UnimplementedSolServer
}

func NewSolServer(svcCtx *svc.ServiceContext) *SolServer {
	return &SolServer{
		svcCtx: svcCtx,
	}
}

func (s *SolServer) Ping(ctx context.Context, in *sol.Request) (*sol.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *SolServer) GetTxByHash(ctx context.Context, in *sol.GetTxByHashRequest) (*sol.Tx, error) {
	l := logic.NewGetTxByHashLogic(ctx, s.svcCtx)
	return l.GetTxByHash(in)
}

func (s *SolServer) Start(ctx context.Context, in *sol.StartRequest) (*sol.Response, error) {
	l := logic.NewStartLogic(ctx, s.svcCtx)
	return l.Start(in)
}

func (s *SolServer) Trade(ctx context.Context, in *sol.TradeRequest) (*sol.Response, error) {
	l := logic.NewTradeLogic(ctx, s.svcCtx)
	return l.Trade(in)
}

func (s *SolServer) GetMarketInfo(ctx context.Context, in *sol.GetMarketInfoRequest) (*sol.MarketInfoResponse, error) {
	l := logic.NewGetMarketInfoLogic(ctx, s.svcCtx)
	return l.GetMarketInfo(in)
}

func (s *SolServer) GetMarketInfoByQuoteMint(ctx context.Context, in *sol.GetMarketInfoByQuoteMintRequest) (*sol.MarketInfoResponse, error) {
	l := logic.NewGetMarketInfoByQuoteMintLogic(ctx, s.svcCtx)
	return l.GetMarketInfoByQuoteMint(in)
}

func (s *SolServer) GetMarketList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.MarketListResponse, error) {
	l := logic.NewGetMarketListLogic(ctx, s.svcCtx)
	return l.GetMarketList(in)
}

func (s *SolServer) GetMarketKlineList(ctx context.Context, in *sol.GetMarketKlineRequest) (*sol.MarketKlineListResponse, error) {
	l := logic.NewGetMarketKlineListLogic(ctx, s.svcCtx)
	return l.GetMarketKlineList(in)
}

func (s *SolServer) GetTrendingMarket(ctx context.Context, in *sol.GetMarketListRequest) (*sol.RealTimeMarketListResponse, error) {
	l := logic.NewGetTrendingMarketLogic(ctx, s.svcCtx)
	return l.GetTrendingMarket(in)
}

func (s *SolServer) GetXTopMarketList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.RealTimeMarketListResponse, error) {
	l := logic.NewGetXTopMarketListLogic(ctx, s.svcCtx)
	return l.GetXTopMarketList(in)
}

func (s *SolServer) GetMemeMarketList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.RealTimeMarketListResponse, error) {
	l := logic.NewGetMemeMarketListLogic(ctx, s.svcCtx)
	return l.GetMemeMarketList(in)
}

func (s *SolServer) GetFollowMarketList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.RealTimeMarketListResponse, error) {
	l := logic.NewGetFollowMarketListLogic(ctx, s.svcCtx)
	return l.GetFollowMarketList(in)
}

func (s *SolServer) SearchCurrencyList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.SearchCurrencyListResponse, error) {
	l := logic.NewSearchCurrencyListLogic(ctx, s.svcCtx)
	return l.SearchCurrencyList(in)
}

func (s *SolServer) MarketActivityList(ctx context.Context, in *sol.GetMarketListRequest) (*sol.MarketActivityListResponse, error) {
	l := logic.NewMarketActivityListLogic(ctx, s.svcCtx)
	return l.MarketActivityList(in)
}

func (s *SolServer) GetMarketDetail(ctx context.Context, in *sol.GetMarketListRequest) (*sol.MarketDetailResponse, error) {
	l := logic.NewGetMarketDetailLogic(ctx, s.svcCtx)
	return l.GetMarketDetail(in)
}
