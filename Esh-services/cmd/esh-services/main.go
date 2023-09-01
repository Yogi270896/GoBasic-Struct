package main

import (
	"Esh-services/internal/app"
	"Esh-services/internal/conf"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var appConfig *conf.Appconfig

func init() {
	appConfig = conf.NewConfig()
}

func main() {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True",
		appConfig.DBConfig.DBUSER,
		appConfig.DBConfig.DBPASSWORD,
		appConfig.DBConfig.DBHOST,
		appConfig.DBConfig.DBPORT,
		appConfig.DBConfig.DBNAME,
	))

	if err != nil {
		log.Println("Error Connecting to the DB", err.Error())
		panic(err)
	}
	db.DB().SetMaxIdleConns(appConfig.DBConfig.MAX_CONNECTIONS)
	db.DB().SetMaxOpenConns(appConfig.DBConfig.MAX_CONNECTIONS)

	defer db.Close()

	//Needs to be config driven
	//db.LogMode(appConfig.Server.PROFILE == "dev" && appConfig.DBConfig.DBLOGMODE)
	appConfig.DB = db
	app.StartApplication(appConfig)
}
