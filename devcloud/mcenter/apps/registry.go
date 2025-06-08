package apps

import (
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/user/api"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/user/impl"

	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/api"
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/impl"

	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/endpoint/impl"

	// 颁发器
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/apps/token/issuers"
	// 鉴权中间件
	_ "122.51.31.227/go-course/go18/devcloud/mcenter/permission"
)
