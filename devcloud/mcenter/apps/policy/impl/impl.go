package impl

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/namespace"
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy"
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/role"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&PolicyServiceImpl{})
}

var _ policy.Service = (*PolicyServiceImpl)(nil)

type PolicyServiceImpl struct {
	ioc.ObjectImpl

	namespace namespace.Service
	role      role.Service
}

func (i *PolicyServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&policy.Policy{})
		if err != nil {
			return err
		}
	}
	i.namespace = namespace.GetService()
	i.role = role.GetService()
	return nil
}

func (i *PolicyServiceImpl) Name() string {
	return policy.AppName
}
