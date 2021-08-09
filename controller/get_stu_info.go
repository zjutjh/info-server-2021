package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zjutjh/info-backend/handler"
	"gorm.io/gorm"
	"net/http"
)

func GetInfo(context *gin.Context) {
	// check from validity
	form, err := handler.BasicCheck(context)
	if err != nil {
		context.JSON(http.StatusOK,
			gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	result, err := handler.QueryInfo(form)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusNotFound, "msg": "RecordNotFound"})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "InternalServerError"})
			return
		}
	}
	// OK
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": result})
}

func GetDorm(context *gin.Context) {
	// check from validity
	form, err := handler.BasicCheck(context)
	if err != nil {
		context.JSON(http.StatusOK,
			gin.H{"status": "fail", "msg": err.Error()})
		return
	}
	// query in database
	result, err := handler.QueryDorm(form)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusNotFound, "msg": "RecordNotFound"})
			return
		} else if errors.Is(err, handler.NotAvailable) {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusServiceUnavailable, "msg": "NotAvailable"})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "InternalServerError"})
			return
		}
	}
	// OK
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": result})
}
