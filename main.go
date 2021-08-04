package main

import (
	"github.com/gin-gonic/gin"
	"info2021/controller"
	"log"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/getStuID", controller.GetID)
		v1.GET("/getInfo", controller.GetInfo)
	}
	if err := router.Run(); err != nil {
		log.Println(err)
	}

}
