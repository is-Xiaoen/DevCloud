package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/label"
	"122.51.31.227/go-course/go18/devcloud/mcenter/test"
)

var (
	svc label.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	svc = label.GetService()
}
