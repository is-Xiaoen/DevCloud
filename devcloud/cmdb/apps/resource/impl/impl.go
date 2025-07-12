package impl

import (
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&ResourceServiceImpl{})
}

var _ resource.Service = (*ResourceServiceImpl)(nil)

type ResourceServiceImpl struct {
	ioc.ObjectImpl
}

func (s *ResourceServiceImpl) Name() string {
	return resource.APP_NAME
}

func (s *ResourceServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&resource.Resource{})
		if err != nil {
			return err
		}
	}
	return nil
}
