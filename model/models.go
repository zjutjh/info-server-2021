package model

type Info struct {
	ID      string `gorm:"primaryKey;comment:身份证号;size:20"`
	UID     string `gorm:"not null;uniqueIndex;comment:学号;size:20"`
	Name    string `gorm:"not null;comment:姓名;size:15"`
	Faculty string `gorm:"comment:学院;size:20"`
	Class   string `gorm:"comment:班级;size:20"`
	Campus  string `gorm:"comment:校区;size:10"`
	House   string `gorm:"comment:寝室楼;size:15"`
	Room    string `gorm:"comment:房间号;size:10"`
	Bed     int8   `gorm:"comment:床号"`
}

type Student struct {
	Info Info `gorm:"embedded;embeddedPrefix:stu_"`
}