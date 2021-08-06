package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"info/handler"
	"net/http"
)

func GetInfo(context *gin.Context) {
	// check from validity
	form, err := handler.BasicCheck(context)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"status": "fail", "msg": err.Error()})
		return
	}
	result, err := handler.QueryInfo(form)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"status": "fail", "msg": "RecordNotFound"})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "msg": "InternalServerError"})
			return
		}
	}
	// OK
	context.JSON(http.StatusOK, gin.H{"status": "ok", "data": result})
}

func GetDorm(context *gin.Context) {
	// check from validity
	form, err := handler.BasicCheck(context)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"status": "fail", "msg": err.Error()})
		return
	}
	// query in database
	result, err := handler.QueryDorm(form)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"status": "fail", "msg": "RecordNotFound"})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "msg": "InternalServerError"})
			return
		}
	}
	// OK
	context.JSON(http.StatusOK, gin.H{"status": "ok", "data": result})
}
