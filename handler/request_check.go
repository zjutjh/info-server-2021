package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"info/model"
	"net/url"
)

//BasicCheck check if 'GET' query valid
func BasicCheck(c *gin.Context) (*model.GetInfo, error) {
	var form model.GetInfo
	err := c.ShouldBind(&form)
	if err != nil || form.Name == "" || form.ID == "" {
		return nil, errors.New("InvalidRequestQuery")
	}
	form.Name, _ = url.QueryUnescape(form.Name)
	return &form, nil
}
