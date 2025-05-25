package impl_test

import (
	"testing"

	"122.51.31.227/go-course/go18/book/v4/apps/comment"
)

func TestAddComment(t *testing.T) {
	ins, err := svc.AddComment(ctx, &comment.AddCommentRequest{
		BookId:  10,
		Comment: "评论测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
