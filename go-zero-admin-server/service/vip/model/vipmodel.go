package model

import (
	"gorm.io/gorm"
)

type (
	VipModel interface {
		GetVipByCardNumber(cardNumber string) (*Vip, error)
		GetVipByCardNumberLike(cardNumber string) ([]Vip, error)
		GetVipById(id uint) (*Vip, error)
		GetVips(page int, limit int, vipS *Vip) ([]Vip, int64, error)
		IsExistVipById(id uint) (bool, error)
		IsExistVipByCardNumber(cardNumber string) (bool, error)
		AddVip(vip *Vip) (bool, error)
		UpdateVip(vip *Vip) (bool, error)
		DelVip(vip *Vip) (bool, error)
		DelSomeVip(ids []uint) (bool, error)
	}

	defaultVipModel struct {
		conn *gorm.DB
	}

	Vip struct {
		gorm.Model
		CardNumber string  `json:"cardNumber" gorm:"type:varchar(191);comment:会员卡号;not null"`     //会员卡号
		Name string `json:"name" gorm:"type:varchar(191);comment:会员名称;not null"`  //会员名称
		Info string `json:"info" gorm:"type:varchar(1000);comment:会员信息;not null"`  //会员信息
	}
)


func NewVipModel(conn *gorm.DB) VipModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Vip{})
	return &defaultVipModel{
		conn: conn,
	}
}


func (d defaultVipModel) GetVipByCardNumber(cardNumber string) (*Vip, error) {
	var vip *Vip
	err := d.conn.Model(&Vip{}).Where("card_number = ?", cardNumber).Find(&vip).Error
	if err != nil {
		return nil, err
	}
	return vip, nil
}

func (d defaultVipModel) GetVipByCardNumberLike(cardNumber string) ([]Vip, error) {
	var vip []Vip
	err := d.conn.Model(&Vip{}).Where("card_number like ?", "%"+cardNumber+"%").Find(&vip).Error
	if err != nil {
		return nil, err
	}
	return vip, nil
}

func (d defaultVipModel) GetVipById(id uint) (*Vip, error) {
	var vip *Vip
	err := d.conn.Model(&Vip{}).Where("id = ?", id).Find(&vip).Error
	if err != nil {
		return nil, err
	}
	return vip, nil
}

func (d defaultVipModel) GetVips(page int, limit int, vipS *Vip) ([]Vip, int64, error) {
	var vips []Vip
	var total int64
	temp := d.conn.Model(&Vip{}).Order("id ASC")
	if vipS.Name != "" {
		temp = temp.Where("name like ?", "%"+vipS.Name+"%")
	}
	if vipS.CardNumber != "" {
		temp = temp.Where("card_number like ?", "%"+vipS.CardNumber+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&vips).Error
	if err != nil {
		return nil, 0, err
	}
	return vips, total, nil
}

func (d defaultVipModel) IsExistVipById(id uint) (bool, error) {
	var total int64
	err := d.conn.Model(&Vip{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (d defaultVipModel) IsExistVipByCardNumber(cardNumber string) (bool, error) {
	var total int64
	err := d.conn.Model(&Vip{}).Where("card_number = ?", cardNumber).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (d defaultVipModel) AddVip(vip *Vip) (bool, error) {
	err := d.conn.Model(vip).Create(vip).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultVipModel) UpdateVip(vip *Vip) (bool, error) {
	err := d.conn.Model(vip).Save(vip).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultVipModel) DelVip(vip *Vip) (bool, error) {
	err := d.conn.Model(vip).Delete(vip).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultVipModel) DelSomeVip(ids []uint) (bool, error) {
	err := d.conn.Model(&Vip{}).Delete(&Vip{}, ids).Error
	if err != nil {
		return false, err
	}
	return true, nil
}