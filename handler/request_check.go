package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"info/model"
)

//BasicCheck check if 'GET' query valid
func BasicCheck(c *gin.Context) (*model.GetInfo, error) {
	var query model.GetInfo
	err := c.ShouldBindQuery(&query)
	if err != nil || query.Name == "" || query.ID == "" {
		return nil, errors.New("InvalidRequestQuery")
	}
	if  len(query.ID) != 18 {
		return nil, errors.New("InvalidIDNumber")
	}
	return &query, nil
}
