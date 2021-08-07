package handler

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/zjutjh/info-backend/data"
	"github.com/zjutjh/info-backend/model"
	"strconv"
	"strings"
)

func ReadInfo(path string, passwd string, sheet string, update bool) {
	InitDB()
	// open file
	var f *excelize.File
	var err error
	if passwd != "" {
		f, err = excelize.OpenFile(path, excelize.Options{Password: passwd})
	} else {
		f, err = excelize.OpenFile(path)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// read
	var rows [][]string
	if sheet != "" {
		rows, err = f.GetRows(sheet)
	} else {
		rows, err = f.GetRows("Sheet1")
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// index of column
	var index = make(map[string]int)
	for i, row := range rows {
		if i == 0 {
			indexSum1 := 0
			indexSum2 := 0
			for j, col := range row {
				if strings.Index(col, "姓名") != -1 {
					index["name"] = j
					indexSum1++
					indexSum2++
				} else if strings.Index(col, "证件号") != -1 {
					index["id"] = j
					indexSum1++
					indexSum2++
				} else if strings.Index(col, "校区") != -1 {
					index["campus"] = j
					indexSum1++
				} else if strings.Index(col, "学院") != -1 {
					index["faculty"] = j
					indexSum1++
				} else if strings.Index(col, "班级") != -1 {
					index["class"] = j
					indexSum1++
				} else if strings.Index(col, "学号") != -1 {
					index["uid"] = j
					indexSum1++
					indexSum2++
				} else if strings.Index(col, "寝室楼") != -1 {
					index["house"] = j
					indexSum2++
				} else if strings.Index(col, "寝室号") != -1 {
					index["room"] = j
					indexSum2++
				} else if strings.Index(col, "床号") != -1 {
					index["bed"] = j
					indexSum2++
				}
			}
			// check validity
			if _, ok := index["uid"]; !update && ok && indexSum1 != 6 {
				fmt.Println("Invalid insert excel sheet")
				return
			}
			if _, ok := index["house"]; ok {
				if indexSum2 != 6 && indexSum2 != 5 {
					fmt.Println("Invalid update excel sheet")
					return
				}
			}
		}
		if i > 0 {
			// insert mode
			if _, ok := index["uid"]; !update && ok {
				stu := model.Student{
					Name:    row[index["name"]],
					ID:      row[index["id"]],
					Campus:  row[index["campus"]],
					Faculty: row[index["faculty"]],
					Class:   row[index["class"]],
					UID:     row[index["uid"]],
				}
				result := data.DB.Model(&model.Student{}).Create(&stu)
				if result.Error != nil {
					fmt.Println(result.Error.Error())
				}
			}
			// update mode
			if _, ok := index["house"]; ok {
				bedNum, err := strconv.Atoi(row[index["bed"]])
				if err != nil {
					fmt.Println(err.Error())
				}
				set := model.Student{
					House: row[index["house"]],
					Room:  row[index["room"]],
					Bed:   int8(bedNum),
				}
				// set "WHERE"
				var where model.Student
				if _, ok := index["id"]; ok {
					where = model.Student{
						Name: row[index["name"]],
						ID:   row[index["id"]],
					}
				} else {
					where = model.Student{
						Name: row[index["name"]],
						UID:  row[index["uid"]],
					}
				}
				result := data.DB.Model(&where).Select("House", "Room", "Bed").Updates(&set)
				if result.Error != nil {
					fmt.Println(result.Error.Error())
				}
			}
		}
	}
}
