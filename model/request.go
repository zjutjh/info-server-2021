package model

//GetInfo request struct of API
type GetInfo struct {
	Name string `json:"stu_name" binding:"required"`
	ID   string `json:"stu_id" binding:"required"`
}
