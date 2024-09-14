package route

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/controller"
	"github.com/PillLabs/plugman-monorepo/metadata/src/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	r.Use(middleware.Cors)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/api/v1", hailing)
	mintGroup := r.Group("/api/v1/mint")
	var mint controller.Mint
	{
		mintGroup.GET("/:address/:mint_type/:count/getSignature", mint.GetSignature)
	}
	orderGroup := r.Group("/api/v1/order")
	var order controller.Order
	{
		orderGroup.POST("/:address", order.PostOrder)
		orderGroup.GET("/:address", order.GetOrders)
	}
	whitelistGroup := r.Group("/api/v1/whitelist")
	var whitelist controller.Whitelist
	{
		whitelistGroup.GET("/:address", whitelist.GetAddressWhitelist)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func hailing(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}
