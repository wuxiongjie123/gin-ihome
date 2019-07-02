package models

import "time"

/* 订单 table_name = ih_order */
type OrderHouse struct {
	OrderId    int       `json:"order_id"`                            //订单编号
	User       []Account   `gorm:"ForeignKey:user_id"`                  //下单的用户编号   //与用户表进行关联
	House      []House   `gorm:"ForeignKey:house_id"`                 //预定的房间编号   //与房屋信息进行关联
	BeginDate  time.Time `gorm:"type:datetime"`                       //预定的起始时间
	EndDate    time.Time `gorm:"type:datetime"`                       //预定的结束时间
	Days       int                                                    //预定总天数
	HousePrice int                                                    //房屋的单价
	Amount     int                                                    //订单总金额
	Status     string    `gorm:"default:'WAIT_ACCEPT'"`               //订单状态
	Comment    string    `gorm:"type:varchar(512)"`                   //订单评论
	Ctime      time.Time `gorm:"auto_now;type:datetime" json:"ctime"` //每次更新此表，都会更新这个字段
	Credit     bool
}
