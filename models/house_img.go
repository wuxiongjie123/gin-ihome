package models

/* 房屋图片 table_name = "ih_house_image"*/
type HouseImage struct {
	HouseImgId int    `json:"house_image_id"`               //图片id
	Url        string `gorm:"type:varchar(256)" json:"url"` //图片url     存放我们房屋的图片
	House      *House `gorm:"ForeignKey:house_id"`          //图片所属房屋编号
}
