package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
	"github.com/zjutjh/info-backend/controller"
	"github.com/zjutjh/info-backend/handler"
	"github.com/zjutjh/info-backend/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	VERSION = "v1.1.3"
)

func main() {
	// set release mode
	gin.SetMode(gin.ReleaseMode)
	// parse cmd args
	var options model.Options
	if _, err := flags.Parse(&options); err != nil {
		return
	}
	// show VERSION
	if options.Version {
		fmt.Println("Info Server version " + VERSION)
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
			os.Exit(2)
		} else {
			fmt.Println(err)
			os.Exit(5)
		}
	}

	// load data from excel file
	if options.LoadData != "" {
		fmt.Println("Start read & load...")
		handler.ReadInfo(options.LoadData, options.Passwd, options.Sheet, options.Update)
		fmt.Println("Load finished")
		return
	}
	// initial database connection
	handler.InitDB()

	// modify routers
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/info", controller.GetInfo)
		v1.POST("/dorm", controller.GetDorm)
		v1.OPTIONS("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
		v1.OPTIONS("/dorm", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	// start server
	var srv *http.Server
	if port := viper.GetString("server-port"); port != "" {
		log.Println("Info Server " + VERSION + " starting at " + port)
		srv = &http.Server{
			Addr:    port,
			Handler: router,
		}
	} else {
		log.Println("Info Server " + VERSION + " starting at :8080")
		srv = &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
