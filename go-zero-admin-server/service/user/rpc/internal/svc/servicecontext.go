package svc

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/user/model"
	"go-zero-admin-server/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	RoleModel model.RoleModel
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
		Config: c,
		UserModel: model.NewUserModel(conn),
		RoleModel: model.NewRoleModel(conn),
	}
}
