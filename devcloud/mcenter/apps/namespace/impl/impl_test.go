package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/namespace"
	"122.51.31.227/go-course/go18/devcloud/mcenter/test"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = namespace.GetService()
}
