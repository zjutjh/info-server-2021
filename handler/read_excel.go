package handler

import (
	"fmt"
	"github.com/spf13/viper"
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

	// find index of column
	var index = make(map[string]int)
	var keys map[string]string
	if !viper.IsSet("excel") {
		keys = make(map[string]string)
		keys["name"] = "姓名"
		keys["id"] = "证件号"
		keys["campus"] = "校区"
		keys["faculty"] = "学院"
		keys["class"] = "班级"
		keys["uid"] = "学号"
		keys["house"] = "寝室楼"
		keys["room"] = "寝室号"
		keys["bed"] = "床号"
	} else {
		keys = viper.GetStringMapString("excel")
	}
	row := rows[0]
	{
		indexSum1 := 0
		indexSum2 := 0
		for j, col := range row {
			if strings.Index(col, keys["name"]) != -1 {
				index["name"] = j
				indexSum1++
				indexSum2++
			} else if strings.Index(col, keys["id"]) != -1 {
				index["id"] = j
				indexSum1++
				indexSum2++
			} else if strings.Index(col, keys["campus"]) != -1 {
				index["campus"] = j
				indexSum1++
			} else if strings.Index(col, keys["faculty"]) != -1 {
				index["faculty"] = j
				indexSum1++
			} else if strings.Index(col, keys["class"]) != -1 {
				index["class"] = j
				indexSum1++
			} else if strings.Index(col, keys["uid"]) != -1 {
				index["uid"] = j
				indexSum1++
				indexSum2++
			} else if strings.Index(col, keys["house"]) != -1 {
				index["house"] = j
				indexSum2++
			} else if strings.Index(col, keys["room"]) != -1 {
				index["room"] = j
				indexSum2++
			} else if strings.Index(col, keys["bed"]) != -1 {
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

	for _, row := range rows[1:] {
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
					//Name: row[index["name"]],
					ID:   row[index["id"]],
				}
			} else {
				where = model.Student{
					//Name: row[index["name"]],
					UID:  row[index["uid"]],
				}
			}
			result := data.DB.Model(&model.Student{}).Where(&where).Select("House", "Room", "Bed").Updates(&set)
			if result.Error != nil {
				fmt.Println(result.Error.Error())
			}
		}
	}
}
