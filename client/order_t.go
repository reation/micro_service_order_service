package main

import (
	"context"
	"github.com/reation/micro_service_order_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	OrderListAddress   = "192.168.1.104:8020"
	OrderDetailAddress = "192.168.1.104:8021"
)

func main() {
	orderList()
	//orderDetail()

}

func orderList() {
	conn, err := grpc.Dial(OrderListAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewOrderListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetOrderList(ctx, &protoc.OrderListRequest{UserId: 1, State: 0, Id: 0, Limit: 20})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("last_id: %d", r.GetLastId())
	log.Println(r.GetOrderList())
	for _, v := range r.OrderList {
		log.Println(v.GetGoodsList())
	}
}

func orderDetail() {
	conn, err := grpc.Dial(OrderDetailAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewOrderDetailClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	orderNum := "2023011623552601234567"
	r, err := c.GetOrderDetail(ctx, &protoc.OrderDetailRequest{OrderNum: orderNum})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.OrderDetail)
}
