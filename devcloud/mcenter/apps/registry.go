package apps

import (
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/user/api"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/user/impl"

	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/api"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/impl"

	// 鉴权
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/endpoint/impl"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/label/impl"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/namespace/impl"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy/impl"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/role/impl"

	// 颁发器
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/issuers"
	// 鉴权中间件
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/permission"
)
