package main

import (
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/proto"
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:2492", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	brandClient = proto.NewGoodsClient(conn) // 确保 brandClient 是 proto.GoodsClient 类型
}

func main() {
	Init()

	//TestGetGoodsDetail()
	//TestGetBrandList()
	//TestGetCategoryBrandList()
	TestGetGoodsList()
	//TestGetGoodsDetail()
	conn.Close()
}
