package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn2() {
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()                                           // 创建默认配置
	consulInfo := global.ServerConfig.ConsulInfo                         // 获取全局服务器配置中的Consul信息
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port) // 格式化地址为 "host:port"

	userSrvHost := ""
	userSrvPort := 0

	client, err := api.NewClient(cfg) // 使用配置创建新的客户端
	if err != nil {                   // 如果创建客户端时发生错误
		panic(err) // 抛出错误
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSerInfo.Name)) // 获取过滤后的服务列表
	//data, err := client.Agent().ServicesWithFilter(`Service == "user-srv"`) // 获取过滤后的服务列表
	if err != nil { // 如果获取服务列表时发生错误
		panic(err) // 抛出错误
	}

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"msg", err.Error(),
		)
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
