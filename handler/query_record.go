package handler

import (
	"errors"
	"info/data"
	"info/model"
)

var NotAvailable = errors.New("no available data yet")

func QueryInfo(form *model.GetInfo) (*model.Info, error) {
	// send SQL query
	var request model.Info
	result := data.DB.Model(&model.Student{}).Where("id = ? AND name = ?", form.ID, form.Name).First(&request)
	// SQL result check
	if result.Error != nil {
		return nil, result.Error
	}
	return &request, nil
}

func QueryDorm(form *model.GetInfo) (*model.Dorm, error) {
	// send SQL query
	var request model.Dorm
	result := data.DB.Model(&model.Student{}).Where("id = ? AND name = ?", form.ID, form.Name).First(&request)
	// SQL result check
	if result.Error != nil {
		return nil, result.Error
	}
	if request.House == "" {
		return nil,NotAvailable
	}
	return &request, nil
}