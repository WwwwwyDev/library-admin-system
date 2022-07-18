package svc

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/systemlog/model"
	"go-zero-admin-server/service/systemlog/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	LoginLogModel model.LoginLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlcfg := c.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlcfg.User,
		mysqlcfg.Password,
		mysqlcfg.Host,
		mysqlcfg.Port,
		mysqlcfg.Name,
	)
	conn := core.NewMysqlGorm(dsn)
	return &ServiceContext{
		Config:        c,
		LoginLogModel: model.NewLoginLogModel(conn),
	}
}
