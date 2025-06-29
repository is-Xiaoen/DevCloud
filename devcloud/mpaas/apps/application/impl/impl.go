package impl

import (
	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Controller().Registry(&ApplicationServiceImpl{})
}

var _ application.Service = (*ApplicationServiceImpl)(nil)

type ApplicationServiceImpl struct {
	ioc.ObjectImpl

	log *zerolog.Logger
}

func (i *ApplicationServiceImpl) Init() error {
	i.log = log.Sub(i.Name())

	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&application.Application{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *ApplicationServiceImpl) Name() string {
	return application.APP_NAME
}
