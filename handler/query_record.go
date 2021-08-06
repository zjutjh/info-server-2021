package handler

import (
	"errors"
	"github.com/zjutjh/info-backend/data"
	"github.com/zjutjh/info-backend/model"
	"gorm.io/gorm"
)

var NotAvailable = errors.New("no available data yet")

func QueryInfo(form *model.GetInfo) (*model.Info, error) {
	// send SQL query
	var request model.Info
	result := data.DB.Model(&model.Student{}).Where(&form).First(&request)
	// SQL result check
	if result.Error != nil {
		return nil, result.Error
	}
	return &request, nil
}

func QueryDorm(form *model.GetInfo) (*model.Dorm, error) {
	// send SQL query
	var request model.Dorm
	result := data.DB.Model(&model.Student{}).Where(&form).First(&request)
	// SQL result check
	if result.Error != nil {
		return nil, result.Error
	}
	if request.House == "" {
		return nil, NotAvailable
	}
	var friends model.Friends
	result = data.DB.Model(&model.Student{}).Where(&model.Student{Campus: request.Campus, House: request.House,
		Room: request.Room}).Order("bed").Find(&friends)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &request, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	request.Friends = friends
	return &request, nil
}
