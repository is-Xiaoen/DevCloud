package impl

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/role"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&RoleServiceImpl{})
}

var _ role.Service = (*RoleServiceImpl)(nil)

type RoleServiceImpl struct {
	ioc.ObjectImpl
}

func (i *RoleServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&role.Role{}, &role.ApiPermission{}, &role.ViewPermission{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *RoleServiceImpl) Name() string {
	return role.AppName
}
