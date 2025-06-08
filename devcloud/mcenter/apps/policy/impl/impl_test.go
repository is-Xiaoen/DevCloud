package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy"
	"122.51.31.227/go-course/go18/devcloud/mcenter/test"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = policy.GetService()
}
