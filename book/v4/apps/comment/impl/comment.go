package impl

import (
	"context"

	"122.51.31.227/go-course/go18/book/v4/apps/book"
	"122.51.31.227/go-course/go18/book/v4/apps/comment"
)

// AddComment implements comment.Service.
func (c *CommentServiceImpl) AddComment(ctx context.Context, in *comment.AddCommentRequest) (*comment.Comment, error) {
	// 能不能 直接Book Service的具体实现
	// (&impl.BookServiceImpl{}).DescribeBook(ctx, nil)
	// 依赖接口，面向接口编程, 不依赖具体实现
	book.GetService().DescribeBook(ctx, nil)
	panic("unimplemented")
}
