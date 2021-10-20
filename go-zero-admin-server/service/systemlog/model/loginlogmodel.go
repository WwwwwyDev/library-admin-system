package model

import (
	"gorm.io/gorm"
)

type (
	LoginLogModel interface {
		ClearLoginLogs() (bool,error)
		GetLoginLogs(page int, limit int, loginLogS *LoginLog) ([]LoginLog, int64, error)
		AddLoginLog(loginLog *LoginLog) (bool, error)
	}

	defaultLoginLogModel struct {
		conn *gorm.DB
	}

	LoginLog struct {
		gorm.Model
		Username string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Info     string `json:"info" gorm:"type:varchar(1000);comment:日志信息;not null"`  //日志信息
	}
)

func NewLoginLogModel(conn *gorm.DB) LoginLogModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(LoginLog{})

	return &defaultLoginLogModel{
		conn: conn,
	}
}


func (d defaultLoginLogModel) ClearLoginLogs() (bool, error) {
	err := d.conn.Where("deleted_at IS NULL").Delete(&LoginLog{}).Error
	if err != nil {
		return false, err
	}
	return true,nil
}

func (d defaultLoginLogModel) GetLoginLogs(page int, limit int, loginLogS *LoginLog) ([]LoginLog, int64, error) {
	var loginLogs []LoginLog
	var total int64
	temp := d.conn.Model(&LoginLog{}).Order("id ASC")
	if  loginLogS.Username != "" {
		temp = temp.Where("username like ?", "%"+ loginLogS.Username+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&loginLogs).Error
	if err != nil {
		return nil, 0, err
	}
	return loginLogs, total, nil
}

func (d defaultLoginLogModel) AddLoginLog(loginLog *LoginLog) (bool, error) {
	err := d.conn.Model(loginLog).Create(loginLog).Error
	if err != nil {
		return false,err
	}
	return true,nil
}