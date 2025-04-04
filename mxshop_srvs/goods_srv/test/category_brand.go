package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"
)

func TestGetCategoryBrandList() {
	rsp, err := brandClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: 135475,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}
