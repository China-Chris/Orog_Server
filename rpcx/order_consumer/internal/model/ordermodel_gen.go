// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderFieldNames          = builder.RawFieldNames(&Order{}, true)
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "id"), ",")
	orderRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(orderFieldNames, "id"))

	cachePublicOrderIdPrefix        = "cache:public:order:id:"
	cachePublicOrderOrderHashPrefix = "cache:public:order:orderHash:"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		FindOneByOrderHash(ctx context.Context, orderHash string) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error
		LoadSolOrders(ctx context.Context) ([]Order, error)
		QueryOrders(ctx context.Context, chainId string, orderType int64, side int64, paymentStatus int64) ([]Order, error)
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id              int64          `db:"id"`
		OrderHash       string         `db:"order_hash"`
		Status          int64          `db:"status"`
		Message         sql.NullString `db:"message"`
		ChainId         string         `db:"chain_id"`
		MarketAddress   string         `db:"market_address"`
		Side            int64          `db:"side"`
		Type            int64          `db:"type"`
		LimitOrderType  int64          `db:"limit_order_type"`
		Price           float64        `db:"price"`
		Amount          float64        `db:"amount"`
		Slippage        float64        `db:"slippage"`
		FilledAmount    float64        `db:"filled_amount"`
		RemainingAmount float64        `db:"remaining_amount"`
		CreatedAt       time.Time      `db:"created_at"`
		UpdatedAt       time.Time      `db:"updated_at"`
		AccountAddress  string         `db:"account_address"`
		PaymentStatus   int64          `db:"payment_status"`
		TransactionHash sql.NullString `db:"transaction_hash"`
		CancelReason    sql.NullString `db:"cancel_reason"`
		OpenMev         bool           `db:"open_mev"`
		RebateStatus    int64          `db:"rebate_status"`
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."order"`,
	}
}

func (m *defaultOrderModel) withSession(session sqlx.Session) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      `"public"."order"`,
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicOrderIdKey := fmt.Sprintf("%s%v", cachePublicOrderIdPrefix, id)
	publicOrderOrderHashKey := fmt.Sprintf("%s%v", cachePublicOrderOrderHashPrefix, data.OrderHash)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicOrderIdKey, publicOrderOrderHashKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	publicOrderIdKey := fmt.Sprintf("%s%v", cachePublicOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, publicOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", orderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) FindOneByOrderHash(ctx context.Context, orderHash string) (*Order, error) {
	publicOrderOrderHashKey := fmt.Sprintf("%s%v", cachePublicOrderOrderHashPrefix, orderHash)
	var resp Order
	err := m.QueryRowIndexCtx(ctx, &resp, publicOrderOrderHashKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where order_hash = $1 limit 1", orderRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, orderHash); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	publicOrderIdKey := fmt.Sprintf("%s%v", cachePublicOrderIdPrefix, data.Id)
	publicOrderOrderHashKey := fmt.Sprintf("%s%v", cachePublicOrderOrderHashPrefix, data.OrderHash)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OrderHash, data.Status, data.Message, data.ChainId, data.MarketAddress, data.Side, data.Type, data.LimitOrderType, data.Price, data.Amount, data.Slippage, data.FilledAmount, data.RemainingAmount, data.CreatedAt, data.UpdatedAt, data.AccountAddress, data.PaymentStatus, data.TransactionHash, data.CancelReason, data.OpenMev, data.RebateStatus)
	}, publicOrderIdKey, publicOrderOrderHashKey)
	return ret, err
}

func (m *defaultOrderModel) LoadSolOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	query := fmt.Sprintf("SELECT %s FROM %s WHERE chain_id = $1", orderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &orders, query, "sol")
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (m *defaultOrderModel) QueryOrders(ctx context.Context, chainId string, orderType int64, side int64, paymentStatus int64) ([]Order, error) {
	var orders []Order
	query := fmt.Sprintf("SELECT %s FROM %s WHERE chain_id = $1 AND type = $2 AND side = $3 AND payment_status = $4", orderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &orders, query, chainId, orderType, side, paymentStatus)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (m *defaultOrderModel) Update(ctx context.Context, newData *Order) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicOrderIdKey := fmt.Sprintf("%s%v", cachePublicOrderIdPrefix, data.Id)
	publicOrderOrderHashKey := fmt.Sprintf("%s%v", cachePublicOrderOrderHashPrefix, data.OrderHash)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.OrderHash, newData.Status, newData.Message, newData.ChainId, newData.MarketAddress, newData.Side, newData.Type, newData.LimitOrderType, newData.Price, newData.Amount, newData.Slippage, newData.FilledAmount, newData.RemainingAmount, newData.CreatedAt, newData.UpdatedAt, newData.AccountAddress, newData.PaymentStatus, newData.TransactionHash, newData.CancelReason, newData.OpenMev, newData.RebateStatus)
	}, publicOrderIdKey, publicOrderOrderHashKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}
