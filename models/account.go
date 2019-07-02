package models

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

// 用户table_name = ih_user
type Account struct {
	UserId       int    `grom:"primary_key" json:"user_id"`            //用户编号
	Name         string `gorm:"type:varchar(32)" json:"name"`          //用户昵称
	PasswordHash string `gorm:"type:varchar(128)" json:"password"`     //用户密码加密的
	Mobile       string `gorm:"type:varchar(11);unique" json:"mobile"` //手机号
	RealName     string `gorm:"type:varchar(32)" json:"real_name"`     //真实姓名  实名认证
	IdCard       string `grom:"type:varchar(20)" json:"id_card"`       //身份证号  实名认证
	AvatarUrl    string `grom:"type:varchar(256)" json:"avatar_url"`   //用户头像路径       通过fastdfs进行图片存储
	Houses       []House                                               //用户发布的房屋信息  一个人多套房
	Order        []OrderHouse                                          //用户下的订单       一个人多次订单
}

func (u *Account) TableName() string {
	return "user"
}
// 创建用户
func (u *Account) CreateAccount() error {
	return db.Create(&u).Error
}

//更新用户
func (u *Account) UpdateAccount() error {
	return db.Update(&u).Error
}

//查询用户
func SelectAccount(mobile string) (*Account, error) {
	var count int
	u := &Account{}
	d := db.Where("mobile = ?", mobile).First(&u)
	d.Count(&count)
	log.Println(d.Count(&count))
	if count <= 0 {
		return u, nil
	}
	return u, d.Error
}

//md5加密
func PassWordMd5(s string) string {
	m := md5.New()
	return hex.EncodeToString(m.Sum([]byte(s)))
}