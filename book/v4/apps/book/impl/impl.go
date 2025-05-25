package impl

import (
	"122.51.31.227/go-course/go18/book/v4/apps/book"
)

// 怎么知道他有没有实现该业务, 可以通过类型约束
// var _ book.Service = &BookServiceImpl{}

//	&BookServiceImpl 的 nil对象
//
// int64(1)  int64 1
// *BookServiceImpl(nil)
var _ book.Service = (*BookServiceImpl)(nil)

// Book业务的具体实现
type BookServiceImpl struct {
}
