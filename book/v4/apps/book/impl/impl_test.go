package impl_test

import (
	"context"

	"122.51.31.227/go-course/go18/book/v4/apps/book"
	"122.51.31.227/go-course/go18/book/v4/test"
)

var ctx = context.Background()
var svc book.Service

func init() {
	test.DevelopmentSet()

	svc = book.GetService()
}
