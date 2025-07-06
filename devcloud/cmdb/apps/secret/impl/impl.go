package impl

import (
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&SecretServiceImpl{})
}

var _ secret.Service = (*SecretServiceImpl)(nil)

type SecretServiceImpl struct {
	ioc.ObjectImpl
}

func (s *SecretServiceImpl) Name() string {
	return secret.APP_NAME
}

func (s *SecretServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&secret.Secret{})
		if err != nil {
			return err
		}
	}
	return nil
}
