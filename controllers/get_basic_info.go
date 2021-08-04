package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetID(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func GetInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "this is Info"})
}