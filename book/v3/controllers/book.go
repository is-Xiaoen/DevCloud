package controllers

import (
	"context"

	"122.51.31.227/go-course/go18/book/v3/config"
	"122.51.31.227/go-course/go18/book/v3/models"
)

var Book = &BookController{}

type BookController struct {
}

func NewGetBookRequest(bookNumber string) *GetBookRequest {
	return &GetBookRequest{
		BookNumber: bookNumber,
	}
}

type GetBookRequest struct {
	BookNumber string
	// RequestId  string
	// ...
}

// 核心功能
// ctx: Trace, 支持请求的取消, request_id
// GetBookRequest 为什么要把他封装为1个对象, GetBook(ctx context.Context, BookNumber string), 保证你的接口的签名的兼容性
// BookController.GetBook(, "")
func (c *BookController) GetBook(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	// context.WithValue(ctx, "request_id", 111)
	// ctx.Value("request_id")

	bookInstance := &models.Book{}
	// 需要从数据库中获取一个对象
	if err := config.DB().Where("id = ?", in.BookNumber).Take(bookInstance).Error; err != nil {
		return nil, err
	}

	return bookInstance, nil
}
