// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderInfoFieldNames          = builder.RawFieldNames(&OrderInfo{})
	orderInfoRows                = strings.Join(orderInfoFieldNames, ",")
	orderInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	orderInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	orderInfoModel interface {
		Insert(ctx context.Context, data *OrderInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderInfo, error)
		Update(ctx context.Context, data *OrderInfo) error
		Delete(ctx context.Context, id int64) error
		GetOrderInfoByOrderNum(ctx context.Context, orderNum string) (*[]OrderInfo, error)
	}

	defaultOrderInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OrderInfo struct {
		Id                 int64     `db:"id"`
		OrderNum           string    `db:"order_num"`             // 订单号
		UserId             int64     `db:"user_id"`               // 用户ID
		GoodsId            int64     `db:"goods_id"`              // 商品ID
		GoodsPrices        float64   `db:"goods_prices"`          // 商品单价
		GoodsNum           int64     `db:"goods_num"`             // 商品数量
		GoodsAllPrices     float64   `db:"goods_all_prices"`      // 商品应有总价=商品单价*商品数量
		GoodsAllRealPrices float64   `db:"goods_all_real_prices"` // 商品实际总价
		OrderPrices        float64   `db:"order_prices"`          // 订单应有价格
		OrderRealPrices    float64   `db:"order_real_prices"`     // 订单实际价格
		PayId              string    `db:"pay_id"`                // 支付号
		BusinessId         int64     `db:"business_id"`           // 商家ID
		State              int64     `db:"state"`                 // 订单状态 1：未支付 11：支付中 21：支付成功 31：支付失败 41：订单超时 51：退单
		UpdateTime         time.Time `db:"update_time"`
		CreateTime         time.Time `db:"create_time"`
	}
)

func newOrderInfoModel(conn sqlx.SqlConn) *defaultOrderInfoModel {
	return &defaultOrderInfoModel{
		conn:  conn,
		table: "`order_info`",
	}
}

func (m *defaultOrderInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderInfoModel) FindOne(ctx context.Context, id int64) (*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderInfoRows, m.table)
	var resp OrderInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) GetOrderInfoByOrderNum(ctx context.Context, orderNum string) (*[]OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where `order_num` = ? ", orderInfoRows, m.table)
	var resp []OrderInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, orderNum)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) Insert(ctx context.Context, data *OrderInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderNum, data.UserId, data.GoodsId, data.GoodsPrices, data.GoodsNum, data.GoodsAllPrices, data.GoodsAllRealPrices, data.OrderPrices, data.OrderRealPrices, data.PayId, data.BusinessId, data.State)
	return ret, err
}

func (m *defaultOrderInfoModel) Update(ctx context.Context, data *OrderInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OrderNum, data.UserId, data.GoodsId, data.GoodsPrices, data.GoodsNum, data.GoodsAllPrices, data.GoodsAllRealPrices, data.OrderPrices, data.OrderRealPrices, data.PayId, data.BusinessId, data.State, data.Id)
	return err
}

func (m *defaultOrderInfoModel) tableName() string {
	return m.table
}