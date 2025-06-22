package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"122.51.31.227/go-course/go18/devcloud/mpaas/test"
)

var (
	svc application.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	svc = application.GetService()
}
