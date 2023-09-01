package services

import (
	"Esh-services/internal/conf"
	"Esh-services/internal/errors"
	"Esh-services/internal/models"
	"Esh-services/internal/repo"
	"fmt"
	"time"
)

type MerchantService struct {
	Appconfig *conf.Appconfig
	Repo      *repo.MerchantRepo
}

func NewMerchantServices(appconf *conf.Appconfig) *MerchantService {
	return &MerchantService{
		Appconfig: appconf,
		Repo:      repo.NewMerchantRepo(appconf),
	}
}

func (oc *MerchantService) CreateMerchant(merchant *models.Merchant, Authorization string) errors.RestAPIError {
	users := models.Merchant{}

	users.Email = merchant.Email
	users.MobileNumber = merchant.MobileNumber
	users.CreatedAt = time.Now()

	// if err := oc.Repo.DB.Find(&merchant).Error; err == nil {
	// 	fmt.Print(err)
	// 	return errors.NewBadRequestError("Email Already Exists")
	// }
	if reserr := oc.Repo.GetByUserLogin(&users); reserr == errors.NO_ERROR() {
		return errors.NewBadRequestError("USER Already Exits")
	}
	if reserr := oc.Repo.Create(&users); reserr != errors.NO_ERROR() {
		return reserr
	}
	fmt.Println("DB RSULT")
	return errors.NO_ERROR()
}

func (oc *MerchantService) GetAllMerchant() {
	fmt.Println("into services")
	oc.Repo.GetAllMerchant()

}
