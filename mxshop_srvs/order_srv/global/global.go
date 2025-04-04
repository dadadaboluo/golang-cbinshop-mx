package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/order_srv/config"
	"mxshop_srvs/order_srv/proto"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)

//
//func init() {
//
//	dsn := "root:hsp@tcp(127.0.0.1:3306)/mxshop_order_srv?charset=utf8mb4&parseTime=True&loc=Local"
//
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold: time.Second,
//			LogLevel:      logger.Info,
//			Colorful:      true,
//		},
//	)
//
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//}
