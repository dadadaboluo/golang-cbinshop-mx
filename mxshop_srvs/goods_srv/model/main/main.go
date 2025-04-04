package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"os"
	"strconv"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {

	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("generic password", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println("1: ", len(newPassword))
	//fmt.Println("2: ", salt)
	//fmt.Println("3: ", encodedPwd)
	//passwordInfo := strings.Split(newPassword, "$")
	//fmt.Println("4: ", passwordInfo)
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check)

	//dsn := "root:hsp@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)
	//
	//// 全局模式
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//_ = db.AutoMigrate(
	//	&model.Category{},
	//	&model.Banner{}, &model.Brands{},
	//	&model.Goods{}, &model.GoodsCategoryBrand{})

	Mysql2Es()
}

func Mysql2Es() {
	dsn := "root:hsp@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	host := "http://192.168.56.10:9200"
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	var goods []model.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err = global.EsClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}

	}
}
