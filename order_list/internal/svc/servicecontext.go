package svc

import (
	"fmt"
	"github.com/reation/micro_service_goods_service/goods_detail/goodsdetail"
	"github.com/reation/micro_service_order_service/model"
	"github.com/reation/micro_service_order_service/order_list/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	OrderInfoModel model.OrderInfoModel
	GoodsService   goodsdetail.GoodsDetail
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.StockTable.User,
		c.Mysql.StockTable.Passwd,
		c.Mysql.StockTable.Addr,
		c.Mysql.StockTable.Port,
		c.Mysql.StockTable.DBName,
	)
	return &ServiceContext{
		Config:         c,
		GoodsService:   goodsdetail.NewGoodsDetail(zrpc.MustNewClient(c.GoodsService)),
		OrderInfoModel: model.NewOrderInfoModel(sqlx.NewMysql(dataSourceUrl)),
	}
}
