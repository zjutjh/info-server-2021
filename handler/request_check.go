package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zjutjh/info-backend/model"
)

//BasicCheck check if 'POST' query valid
func BasicCheck(c *gin.Context) (*model.GetInfo, error) {
	var form model.GetInfo
	err := c.ShouldBindJSON(&form)
	if err != nil || form.Name == "" || form.ID == "" {
		return nil, errors.New("InvalidRequestQuery")
	}
	return &form, nil
}
