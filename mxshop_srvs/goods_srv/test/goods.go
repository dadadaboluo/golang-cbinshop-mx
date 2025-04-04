package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"
)

func TestGetGoodsList() {
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 130361,
		PriceMin:    90,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(rsp.Total)
	fmt.Println("33", rsp.Data)
	//for _, good := range rsp.Data {
	//	fmt.Println(good.Name, good.ShopPrice)
	//}
}

func TestBatchGetGoods() {
	rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{421, 422, 423},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Total)

	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}
func TestGetGoodsDetail() {
	rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)
}
