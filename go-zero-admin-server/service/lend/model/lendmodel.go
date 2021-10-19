package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	LendModel interface {
		IsLend(bookId uint64) (bool, error)
		GetLends(page int, limit int, lendS *Lend) ([]Lend, int64, error)
		AddLend(lend *Lend) (bool, error)
		DelLend(lend *Lend) (bool, error)
		IsExistLendById(id uint64) (bool,error)
	}

	defaultLendModel struct {
		conn *gorm.DB
	}

	Lend struct {
		gorm.Model
		VipCardNumber string    `json:"vipCardNumber" gorm:"type:varchar(191);comment:借阅会员;not null"` //借阅会员
		VipId         uint      `json:"vipId" gorm:"comment:借阅会员id;not null"`                         //借阅会员Id
		BookName      string      `json:"bookName"  gorm:"type:varchar(191);comment:借阅图书;not null"`     //借阅图书
		BookId        uint      `json:"bookId" gorm:"comment:图书id;not null"`                          //图书Id
		LendTime      time.Time `json:"lendTime" gorm:"comment:借阅时间;not null"`                        //借阅时间
	}
)


func NewLendModel(conn *gorm.DB) LendModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(Lend{})

	return &defaultLendModel{
		conn: conn,
	}
}


func (d defaultLendModel) GetLends(page int, limit int, lendS *Lend) ([]Lend, int64, error) {
	var lends []Lend
	var total int64
	temp := d.conn.Model(&Lend{}).Order("id ASC")
	if lendS.VipCardNumber != "" {
		temp = temp.Where("vip_card_number like ?", "%"+lendS.VipCardNumber+"%")
	}
	if lendS.BookName != "" {
		temp = temp.Where("book_name like ?", "%"+lendS.BookName+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&lends).Error
	if err != nil {
		return nil, 0, err
	}
	return lends, total, nil
}

func (d defaultLendModel) AddLend(lend *Lend) (bool, error) {
	err := d.conn.Model(lend).Create(lend).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultLendModel) DelLend(lend *Lend) (bool, error) {
	err := d.conn.Model(lend).Delete(lend).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultLendModel) IsLend(bookId uint64) (bool, error) {
	var lend []Lend
	err := d.conn.Where("book_id=?", bookId).Find(&lend).Error
	if err != nil {
		return false, err
	}
	if len(lend) > 0 {
		return true, err
	}
	return false, nil
}

func (d defaultLendModel) IsExistLendById(id uint64) (bool, error) {
	var total int64
	err := d.conn.Model(&Lend{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}