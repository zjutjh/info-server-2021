package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"info/controller"
	"info/model"
)

var Options struct {
	ConfigPath string `short:"c" long:"config" description:"[PATH] Set config path"`
}

func main() {
	// parse cmd args
	if _, err := flags.Parse(&Options); err != nil {
		panic(err)
	}
	if Options.ConfigPath != "" {
		viper.SetConfigFile(Options.ConfigPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/info2021")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file 'config.yml' not found")
		} else {
			panic(err)
		}
	}

	// initial database connection
	db := initDB()
	if db == nil {
		// exit when connecting fail
		return
	}

	// modify routers
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/getStuID", controller.GetID)
		v1.GET("/getInfo", controller.GetInfo)
	}

	// start server
	if err := router.Run(); err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	user := viper.GetString("username")
	passwd := viper.GetString("password")
	database := viper.GetString("database")
	hostname := viper.GetString("hostname")
	port := viper.GetString("port")
	if user == "" || passwd == "" || database == "" {
		panic("Invalid config")
	}
	var dsn string
	if hostname == "" || port == "" {
		dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, database)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, hostname, port, database)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&model.StuInfo{}); err != nil {
		panic(err)
	}
	return db
}
