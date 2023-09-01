package conf

import (
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Appconfig struct {
	DBConfig dbConfig
	DB       *gorm.DB
}

type dbConfig struct {
	DBUSER          string
	DBPASSWORD      string
	DBHOST          string
	DBNAME          string
	DBPORT          string
	DBLOGMODE       bool
	MAX_CONNECTIONS int
}

func NewConfig() *Appconfig {
	appconifg := Appconfig{
		DBConfig: dbConfig{
			DBUSER:          getEnv("DBUSER", "root"),
			DBPASSWORD:      getEnv("DBPASSWORD", ""),
			DBNAME:          getEnv("DBNAME", "ymerchantdb"),
			DBHOST:          getEnv("DBHOST", "localhost"),
			DBPORT:          getEnv("DBPORT", "3306"),
			DBLOGMODE:       getEnvAsBool("DBLOGMODE", false),
			MAX_CONNECTIONS: getEnvAsInt("MAX_CONNECTIONS", 10),
		},
	}

	return &appconifg
}

func getEnv(key string, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		//log.Printf("Got value: %s for the key: %s \n", key, value)
		return value
	}
	return defaultval
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := getEnv(name, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	return defaultVal
}
