package impl

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func init() {
	ioc.Controller().Registry(&UserServiceImpl{})
}

var _ user.Service = (*UserServiceImpl)(nil)

// 他是user service 服务的控制器
type UserServiceImpl struct {
	ioc.ObjectImpl
}

func (i *UserServiceImpl) Init() error {
	// 自动创建表
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&user.User{})
		if err != nil {
			return err
		}
	}
	return nil
}

// 定义托管到Ioc里面的名称
func (i *UserServiceImpl) Name() string {
	return user.APP_NAME
}
