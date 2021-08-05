package model

type GetInfo struct {
	Name string `form:"stu_name" binding:"required"`
	ID string `form:"stu_id" binding:"required"`
}
