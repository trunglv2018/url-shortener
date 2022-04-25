package config

import (
	"log"
	"os"
	"url-shortener/helpers/db"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper
var logr = logrus.New()

func init() {
	initConfig()
	initDB()
}
func initConfig() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath(".")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		os.Exit(0)
	}
	logr.Out = os.Stdout

	file, err := os.OpenFile("iam-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logr.Out = file
	} else {
		logr.Info("Failed to log to file, using default stderr")
	}
}

func GetConfig() *viper.Viper {
	return config
}

func initDB() {
	var endpoint = GetConfig().GetString("arrango_db.endpoint")
	var uname = GetConfig().GetString("arrango_db.uname")
	var password = GetConfig().GetString("arrango_db.password")
	db.ConnectDB(endpoint, uname, password)
}
