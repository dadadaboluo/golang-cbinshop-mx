package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"mxshop_srvs/goods_srv/proto"
	_ "mxshop_srvs/goods_srv/proto"
)

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, category := range rsp.Data {
		fmt.Println(category.Name)
	}
}
func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 130358,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.SubCategorys)
}
