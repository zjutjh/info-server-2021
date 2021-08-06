package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
	"info/controller"
	"info/handler"
	"info/model"
)

func main() {
	// set release mode
	gin.SetMode(gin.ReleaseMode)
	// parse cmd args
	var options model.Options
	if _, err := flags.Parse(&options); err != nil {
		return
	}

	// read config from file(yaml）
	if options.ConfigPath != "" {
		viper.SetConfigFile(options.ConfigPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/info")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file 'config.yml' not found")
			return
		} else {
			fmt.Println(err)
			return
		}
	}

	// load data from excel file
	if options.LoadData != "" {
		handler.ReaInfo(options.LoadData, options.Passwd)
		fmt.Println("Load finished")
		return
	}
	// initial database connection
	handler.InitDB()

	// modify routers
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/info", controller.GetInfo)
		v1.GET("/dorm", controller.GetDorm)
	}

	// start server
	var err error
	if port := viper.GetString("server-port"); port != "" {
		err = router.Run(port)
		fmt.Println("Info server started at " + port)
	} else {
		err = router.Run(":8080")
		fmt.Println("Info server started at :8080")
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
