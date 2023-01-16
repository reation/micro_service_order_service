package logic

import (
	"context"
	"github.com/reation/micro_service_goods_service/goods_detail/goodsdetail"
	"github.com/reation/micro_service_order_service/config"
	"github.com/reation/micro_service_order_service/model"
	"github.com/reation/micro_service_order_service/order_detail/internal/svc"
	"github.com/reation/micro_service_order_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *protoc.OrderDetailRequest) (*protoc.OrderDetailResponse, error) {
	orderInfo, err := l.svcCtx.OrderInfoModel.GetOrderInfoByOrderNum(l.ctx, in.OrderNum)
	if err == model.ErrNotFound {
		return &protoc.OrderDetailResponse{States: config.RETURN_STATES_EMPTY, OrderDetail: nil}, nil
	}
	if err != nil {
		return &protoc.OrderDetailResponse{States: config.RETURN_STATES_ERROR, OrderDetail: nil}, nil
	}
	var goodsIDList = make([]*goodsdetail.GoodsDetailGoodsIDList, len(*orderInfo))
	for k, v := range *orderInfo {
		goodsIDList[k] = &goodsdetail.GoodsDetailGoodsIDList{Id: v.Id}
	}
	states, resp := l.GetOrderGoodsList(goodsIDList, orderInfo)
	return &protoc.OrderDetailResponse{States: states, OrderDetail: resp}, nil
}

func (l *GetOrderDetailLogic) GetOrderGoodsList(idList []*goodsdetail.GoodsDetailGoodsIDList, orderInfo *[]model.OrderInfo) (int64, *protoc.OrderDetailOrderInfo) {
	goodsListResponse, err := l.svcCtx.GoodsService.GetGoodsListByIDList(l.ctx, &goodsdetail.GetGoodsListByIDListRequest{IdList: idList})
	if err != nil {
		return config.RETURN_STATES_ERROR, nil
	}

	if goodsListResponse.GetStates() != config.RETURN_STATES_NORMAL {
		return goodsListResponse.GetStates(), nil
	}

	var goodsInfoMap = make(map[int64]goodsdetail.GoodsDataInfo)
	for _, v := range goodsListResponse.GetGoodList() {
		goodsInfoMap[v.GetId()] = *v
	}

	var orderNum string
	var orderPrices float64
	var orderRealPrices float64
	var state int64
	var updateTime string
	var createTime string
	var respGoodsInfo = make([]*protoc.OrderDetailOrderGoodsData, len(idList))
	for k, v := range *orderInfo {
		respGoodsInfo[k] = &protoc.OrderDetailOrderGoodsData{
			GoodsId:            v.GoodsId,
			Title:              goodsInfoMap[v.GoodsId].Title,
			Cover:              goodsInfoMap[v.GoodsId].Cover,
			BusinessId:         goodsInfoMap[v.GoodsId].BusinessId,
			BusinessName:       goodsInfoMap[v.GoodsId].BusinessName,
			GoodsPrices:        v.GoodsPrices,
			GoodsNum:           v.GoodsNum,
			GoodsAllPrices:     v.GoodsAllPrices,
			GoodsAllRealPrices: v.GoodsAllRealPrices,
			UpdateTime:         v.UpdateTime.Format("2006-01-02 15:04:05"),
			CreateTime:         v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		orderNum = v.OrderNum
		orderPrices = v.OrderPrices
		orderRealPrices = v.OrderRealPrices
		state = v.State
		updateTime = v.UpdateTime.Format("2006-01-02 15:04:05")
		createTime = v.CreateTime.Format("2006-01-02 15:04:05")
	}

	resp := &protoc.OrderDetailOrderInfo{
		OrderNum:        orderNum,
		OrderPrices:     orderPrices,
		OrderRealPrices: orderRealPrices,
		State:           state,
		GoodsDetail:     respGoodsInfo,
		UpdateTime:      updateTime,
		CreateTime:      createTime,
	}

	return goodsListResponse.GetStates(), resp
}
