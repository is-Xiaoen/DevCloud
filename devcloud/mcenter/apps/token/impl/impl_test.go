package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/token"
	"122.51.31.227/go-course/go18/devcloud/mcenter/test"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	svc = token.GetService()
}
