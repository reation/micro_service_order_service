package logic

import (
	"context"

	"github.com/reation/micro_service_order_service/order_list/internal/svc"
	"github.com/reation/micro_service_order_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *protoc.OrderListRequest) (*protoc.OrderListResponse, error) {
	// todo: add your logic here and delete this line

	return &protoc.OrderListResponse{}, nil
}
