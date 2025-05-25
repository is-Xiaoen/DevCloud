package apps

// 业务加载区, 选择性的价值的业务处理对象

import (
	// Api Impl
	_ "122.51.31.227/go-course/go18/book/v4/apps/book/api"

	// Service Impl
	_ "122.51.31.227/go-course/go18/book/v4/apps/book/impl"
	_ "122.51.31.227/go-course/go18/book/v4/apps/comment/impl"
)
