package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"info/data"
	"info/handler"
	"info/model"
	"net/http"
)

func GetID(context *gin.Context) {
	query, err := handler.BasicCheck(context)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"status": "fail", "msg": err.Error()})
		return
	}

	// send SQL query
	var stu model.Student
	result := data.DB.Where("stu_id = ? AND stu_name = ?", query.ID, query.Name).First(&stu)
	// SQL result check
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound,
			gin.H{"status": "fail", "msg": "RecordNotFound"})
		return
	} else if result.Error != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"status": "fail", "msg": "InternalServerError"})
		return
	}
	// OK
	context.JSON(http.StatusOK, gin.H{"status": "ok", "stu_id": stu.Info.ID})
}

func GetMoreInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "this is Info"})
}
