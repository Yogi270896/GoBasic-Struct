package repo

import (
	"Esh-services/internal/conf"
	"Esh-services/internal/errors"
	"Esh-services/internal/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type MerchantRepo struct {
	DB        *gorm.DB
	Appconfig *conf.Appconfig
}

func NewMerchantRepo(appconf *conf.Appconfig) *MerchantRepo {
	return &MerchantRepo{
		Appconfig: appconf,
		DB:        appconf.DB,
	}
}

func (s *MerchantRepo) Create(user *models.Merchant) errors.RestAPIError {
	if err := s.DB.Table(models.MerchantTable).Save(user).Error; err != nil {
		fmt.Println(err)
		fmt.Println("lllll")
		return errors.NewInternalServerError(fmt.Sprintf("error while creating User: %+v", err.Error()))
	}

	return errors.NO_ERROR()

}

func (s *MerchantRepo) GetByUserLogin(user *models.Merchant) errors.RestAPIError {

	if err := s.DB.Table(models.MerchantTable).Where("email=?", user.Email).Find(user).Error; err != nil {
		fmt.Print(err)
		return errors.NewNotFoundError(fmt.Sprintf("Could not find any User : %+v", err.Error()))
	}
	return errors.NO_ERROR()
}

func (s *MerchantRepo) GetAllMerchant() {
	fmt.Println("INto all repo")
	var mctemp []models.Merchant
	err := s.DB.Table(models.MerchantTable).Find(&mctemp)
	if err.Error != nil {
		fmt.Println(&mctemp, errors.NewInternalServerError("No Record Found"))
	}
	fmt.Println(err.Value)

	//return res, errors.NO_ERROR()
}
