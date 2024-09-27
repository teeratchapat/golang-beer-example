package router

import (
	controller "golang-beer-example/modules/controllers"

	"github.com/gin-gonic/gin"
)

func BeerHandlerRoute(router *gin.RouterGroup) {
	router.GET("/beer/get", controller.BeerGetController)          //hint : /api/beer/get
	router.POST("/beer/insert", controller.BeerPostController)     //hint : /api/beer/insert
	router.PUT("/beer/update", controller.BeerPutController)       //hint : /api/beer/update
	router.DELETE("/beer/delete", controller.BeerDeleteController) //hint : /api/beer/delete
}
