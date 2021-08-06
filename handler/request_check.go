package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"info/model"
	"net/url"
)

//BasicCheck check if 'GET' query valid
func BasicCheck(c *gin.Context) (*model.GetInfo, error) {
	var query model.GetInfo
	err := c.ShouldBindQuery(&query)
	if err != nil || query.Name == "" || query.ID == "" {
		return nil, errors.New("InvalidRequestQuery")
	}
	query.Name, _ = url.QueryUnescape(query.Name)
	return &query, nil
}
