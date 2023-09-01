package controllers

import (
	"Esh-services/internal/conf"
	"Esh-services/internal/errors"
	"Esh-services/internal/models"
	"Esh-services/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	Appconfig       *conf.Appconfig
	MerchantService *services.MerchantService
}

func NewMerchantController(appconf *conf.Appconfig) *MerchantController {
	return &MerchantController{
		Appconfig:       appconf,
		MerchantService: services.NewMerchantServices(appconf),
	}
}

func (oc *MerchantController) GetAllMerchant(c *gin.Context) {
	m := "RegisterMerchant"
	fmt.Println(m)

	oc.MerchantService.GetAllMerchant()

	return
}

func (mc *MerchantController) CreateMerchant(c *gin.Context) {
	var merchant models.Merchant
	m := "CreateMerchant"
	fmt.Println(m, "Start")

	if err := c.ShouldBind(&merchant); err != nil {
		//restErr := strings.Split(err.Error(), "Error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resErr := mc.MerchantService.CreateMerchant(&merchant,
		c.Request.Header.Get("Authorization"))
	if resErr != errors.NO_ERROR() {
		c.JSON(401, resErr)
		return
	}
	c.JSON(http.StatusAccepted, "Sucessfully Created")
}
