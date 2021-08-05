package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"info2021/controllers"
	"info2021/model"
	"log"
)

func main() {
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
		v1.GET("/getStuID", controllers.GetID)
		v1.GET("/getInfo", controllers.GetInfo)
	}

	// start server
	if err := router.Run(); err != nil {
		log.Println(err)
	}

}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:passwd@/info2021?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	if err:= db.AutoMigrate(&model.StuInfo{}); err != nil {
		log.Println(err)
	}
	return db
}