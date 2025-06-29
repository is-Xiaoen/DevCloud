package main

import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// mcenter 业务对象
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps"

	// audit 业务对象
	_ "122.51.31.227/go-course/go18/devcloud/audit/apps"

	// mpaas 应用发布
	_ "122.51.31.227/go-course/go18/devcloud/mpaas/apps"

	// 非功能性模块
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
)

func main() {
	// 启动
	cmd.Start()
}
