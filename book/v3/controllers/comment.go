package controllers

import (
	"context"
	"fmt"

	"122.51.31.227/go-course/go18/book/v3/models"
)

var Comment = &CommentController{}

type CommentController struct {
}

type AddCommentRequest struct {
	BookNumber int
}

func (c *CommentController) AddComment(ctx context.Context, in *AddCommentRequest) (*models.Comment, error) {
	// 业务处理的细节
	// 多个业务模块 进行交互
	book, err := Book.GetBook(ctx, NewGetBookRequest(in.BookNumber))

	// if exception.IsApiException(err, exception.CODE_NOT_FOUND) {

	// }
	if err != nil {
		// 获取查询不到报错
		return nil, err
	}
	// 判断book的状态
	fmt.Println(book)
	return nil, nil
}
