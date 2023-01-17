package logic

import (
	"context"
	"github.com/reation/micro_service_goods_service/config"
	"github.com/reation/micro_service_goods_service/goods_detail/goodsdetail"
	"github.com/reation/micro_service_order_service/model"

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

	states, orderNumList := l.getOrderNumList(in.GetUserId(), in.GetState(), in.GetId(), in.GetLimit())
	if states != config.RETURE_STATES_NORMAL {
		return &protoc.OrderListResponse{States: states, OrderList: nil, LastId: in.GetId()}, nil
	}

	orderInfoListMap, goodsIDList, lastID := l.getOrderInfoListByOrderNumList(orderNumList)
	if orderInfoListMap == nil || goodsIDList == nil {
		return &protoc.OrderListResponse{States: config.RETURE_STATES_ERROR, OrderList: nil, LastId: in.GetId()}, nil
	}

	goodsList := l.getGoodsListByGoodsIDList(goodsIDList)
	if goodsList == nil {
		return &protoc.OrderListResponse{States: config.RETURE_STATES_ERROR, OrderList: nil, LastId: in.GetId()}, nil
	}

	var resp = make([]*protoc.OrderListOrderInfo, len(*orderNumList))
	for k, v := range *orderNumList {
		var goodsLists = make([]*protoc.OrderListGoodsInfo, len(orderInfoListMap[v]))
		for m, n := range orderInfoListMap[v] {
			goodsLists[m] = &protoc.OrderListGoodsInfo{
				Id:                 goodsList[n.GoodsId].Id,
				Title:              goodsList[n.GoodsId].Title,
				Cover:              goodsList[n.GoodsId].Cover,
				BusinessId:         goodsList[n.GoodsId].BusinessId,
				BusinessName:       "嘟嘟嘟超市",
				GoodsPrices:        n.GoodsPrices,
				GoodsNum:           n.GoodsNum,
				GoodsAllPrices:     n.GoodsAllPrices,
				GoodsAllRealPrices: n.GoodsAllRealPrices,
			}
		}

		resp[k] = &protoc.OrderListOrderInfo{
			OrderNum:        orderInfoListMap[v][0].OrderNum,
			GoodsList:       goodsLists,
			OrderPrices:     orderInfoListMap[v][0].OrderPrices,
			OrderRealPrices: orderInfoListMap[v][0].OrderRealPrices,
			State:           orderInfoListMap[v][0].State,
			UpdateTime:      orderInfoListMap[v][0].UpdateTime.Format("2006-01-02 15:04:05"),
			CreateTime:      orderInfoListMap[v][0].UpdateTime.Format("2006-01-02 15:04:05"),
		}

	}

	return &protoc.OrderListResponse{States: config.RETURE_STATES_NORMAL, OrderList: resp, LastId: lastID}, nil
}

func (l *GetOrderListLogic) getOrderNumList(userID, state, id, limit int64) (states int64, orderNumList *[]string) {
	if id < 1 {
		id = 0
	}
	if limit < 1 || limit > 30 {
		limit = 20
	}
	orderNumList, err := l.svcCtx.OrderInfoModel.GetOrderNumByUserID(l.ctx, userID, state, id, limit)
	if err == model.ErrNotFound {
		return config.RETURE_STATES_EMPTY, nil
	}
	if err != nil {
		return config.RETURE_STATES_ERROR, nil
	}

	return config.RETURE_STATES_NORMAL, orderNumList
}

func (l *GetOrderListLogic) getOrderInfoListByOrderNumList(orderNumList *[]string) (orderListMap map[string][]model.OrderInfo, GoodsIDList []*goodsdetail.GoodsDetailGoodsIDList, lastID int64) {
	orderInfoList, err := l.svcCtx.OrderInfoModel.GetOrderInfoListByOrderNumList(l.ctx, orderNumList)
	if err != nil {
		return nil, nil, 0
	}
	var orderInfoListMap = make(map[string][]model.OrderInfo)
	var goodsIDList = make([]*goodsdetail.GoodsDetailGoodsIDList, len(*orderInfoList))
	for k, v := range *orderInfoList {
		orderInfoListMap[v.OrderNum] = append(orderInfoListMap[v.OrderNum], v)
		goodsIDList[k] = &goodsdetail.GoodsDetailGoodsIDList{Id: v.GoodsId}
		lastID = v.Id
	}

	return orderInfoListMap, goodsIDList, lastID
}

func (l *GetOrderListLogic) getGoodsListByGoodsIDList(idList []*goodsdetail.GoodsDetailGoodsIDList) map[int64]goodsdetail.GoodsDataInfo {
	goodsListResponse, err := l.svcCtx.GoodsService.GetGoodsListByIDList(l.ctx, &goodsdetail.GetGoodsListByIDListRequest{IdList: idList})
	if err != nil {
		return nil
	}

	if goodsListResponse.GetStates() != config.RETURE_STATES_NORMAL {
		return nil
	}

	var goodsInfoMap = make(map[int64]goodsdetail.GoodsDataInfo)
	for _, v := range goodsListResponse.GetGoodList() {
		goodsInfoMap[v.GetId()] = *v
	}

	return goodsInfoMap
}
