package impl_test

import (
	"context"
	"os"

	"122.51.31.227/go-course/go18/book/v4/apps/comment/impl"
	"github.com/infraboard/mcube/v2/ioc"
)

var ctx = context.Background()
var svc = impl.CommentServiceImpl{}

func init() {
	// import 后自动执行的逻辑
	// 工具对象的初始化, 需要的是绝对路径
	ioc.DevelopmentSetupWithPath(os.Getenv("workspaceFolder") + "/book/v4/application.toml")
}
