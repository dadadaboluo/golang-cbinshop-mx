package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/goods-web/api/goods"
	"mxshop-api/goods-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	zap.S().Info("配置用户相关的url")

	{
		GoodsRouter.GET("", goods.List)
		//GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New) // 需要管理员权限
		GoodsRouter.POST("", goods.New)
		GoodsRouter.GET("/:id", goods.Detail) //获取商品的详情
		//GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		GoodsRouter.DELETE("/:id", goods.Delete)     //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks) //获取商品的库存

		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
	}
}
