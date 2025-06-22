package impl

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&LabelServiceImpl{})
}

var _ label.Service = (*LabelServiceImpl)(nil)

type LabelServiceImpl struct {
	ioc.ObjectImpl
}

func (i *LabelServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&label.Label{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *LabelServiceImpl) Name() string {
	return label.APP_NAME
}
