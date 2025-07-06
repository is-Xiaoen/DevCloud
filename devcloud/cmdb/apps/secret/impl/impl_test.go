package impl_test

import (
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
	"122.51.31.227/go-course/go18/devcloud/cmdb/test"
)

var (
	svc secret.Service
)

func init() {
	test.DevelopmentSetUp()
	svc = secret.GetService()
}
