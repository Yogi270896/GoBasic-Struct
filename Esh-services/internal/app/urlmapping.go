package app

import (
	"Esh-services/internal/conf"
	"Esh-services/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mapurl(appConfig *conf.Appconfig) {

	merchantController := controllers.NewMerchantController(appConfig)
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	api := router.Group("/v1/")
	{
		api.GET("/all", merchantController.GetAllMerchant)
		api.POST("/create", merchantController.CreateMerchant)
	}
}
