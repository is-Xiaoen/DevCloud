package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/role"
	"122.51.31.227/go-course/go18/devcloud/mcenter/test"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = role.GetService()
}
