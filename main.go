package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
	"info/controller"
	"info/handler"
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
	handler.InitDB()

	// modify routers
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/StuID", controller.GetID)
		v1.GET("/Info", controller.GetMoreInfo)
	}

	// start server
	if err := router.Run(); err != nil {
		panic(err)
	}
}


