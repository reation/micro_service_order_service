package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	GoodsService zrpc.RpcClientConf
	Mysql        struct {
		StockTable struct {
			User   string
			Passwd string
			Addr   string
			DBName string
			Port   string
		}
	}
}
