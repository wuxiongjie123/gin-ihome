package models

/* 设施信息 table_name = ih_facility*/
type Facility struct {
	FId    int     `json:"fid"`                //设施编号
	Name   string  `gorm:"type:varchar(32)"`   //设施名字
	Houses []House `gorm:"many2many:ih_house"` //都有哪些房屋有此设施  与房屋表进行关联的
}
