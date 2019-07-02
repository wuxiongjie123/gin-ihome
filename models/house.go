package models

import "time"

// 房屋信息 table_name = ih_house
type House struct {
	HouseId       int          `json:"house_id"`                                           //房屋编号
	User          []Account       `gorm:"ForeignKey:user_id"`                                 //房屋主人的用户编号  与用户进行关联
	Area          []Area       `grom:"ForeignKey:area_id"`                                 //归属地的区域编号   和地区表进行关联
	Title         string       `grom:"type:varchar(64)" json:"title"`                      //房屋标题
	Price         int          `grom:"default:0" json:"price"`                             //单价,单位:分   每次的价格要乘以100
	Address       string       `grom:"type:varchar(512);default:''" json:"address"`        //地址
	RoomCount     int          `grom:"default:1" json:"room_count"`                        //房间数目
	Acreage       int          `grom:"default:0" json:"acreage"`                           //房屋总面积
	Unit          string       `grom:"type:varchar(32);default:''" json:"unit"`            //房屋单元,如 几室几厅
	Capacity      int          `grom:"default:1" json:"capacity"`                          //房屋容纳的总人数
	Beds          string       `grom:"default:''" json:"beds"`                             //房屋床铺的配置
	Deposit       int          `gorm:"default:0" json:"deposit"`                           //押金
	MinDays       int          `gorm:"default:1" json:"min_days"`                          //最少入住的天数
	MaxDays       int          `gorm:"default:0" json:"max_days"`                          //最多入住的天数 0表示不限制
	OrderCount    int          `gorm:"default:0" json:"order_count"`                       //预定完成的该房屋的订单数
	IndexImageUrl string       `gorm:"type:varchar(256);default:''"json:"index_image_url"` //房屋主图片路径
	Facilities    []Facility   `gorm:"many2many:ih_facilities"`                            //房屋设施   与设施表进行关联
	Images        []HouseImage `gorm:"many2many:ih_img_urls"`                              //房屋的图片   除主要图片之外的其他图片地址
	Orders        []OrderHouse `gorm:"many2many:ih_orders"`                                //房屋的订单    与房屋表进行管理
	Ctime         time.Time    `gorm:"auto_now_add;type:datetime" json:"ctime"`
}
