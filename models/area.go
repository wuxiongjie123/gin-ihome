package models

/* 区域信息 table_name = ih_area */ //区域信息是需要我们手动添加到数据库中的
type Area struct {
	AreaId int     `json:"area_id"`                                //区域编号    1	  2	 3
	Name   string  `gorm:"type:varchar(32)" json:"aname"`      //区域名字    海淀 昌平
	House  []House `gorm:"many2many:ih_house" json:"house_id"` //区域所有的房屋   与房屋表进行关联
}
