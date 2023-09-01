package app

import (
	"Esh-services/internal/conf"
	"Esh-services/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication(appconfig *conf.Appconfig) {

	Mapurl(appconfig)
	runMigrateDB(appconfig)
	router.Use(cors.Default())
	router.Run(":8081")
}

func runMigrateDB(appconf *conf.Appconfig) {
	appconf.DB.Table(models.MerchantTable).AutoMigrate(&models.Merchant{})
}
