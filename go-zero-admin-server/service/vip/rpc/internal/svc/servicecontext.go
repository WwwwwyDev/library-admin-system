package svc

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/vip/model"
	"go-zero-admin-server/service/vip/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	VipModel model.VipModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	psqlcfg := c.Postgresql
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		psqlcfg.User,
		psqlcfg.Password,
		psqlcfg.Name,
		psqlcfg.Host,
		psqlcfg.Port,
	)
	conn := core.NewPostgresqlGorm(dsn)
	return &ServiceContext{
		Config: c,
		VipModel: model.NewVipModel(conn),
	}
}
