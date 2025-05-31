package token_test

import (
	"testing"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/token"
)

func TestMakeBearer(t *testing.T) {
	t.Log(token.MakeBearer(24))
	t.Log(token.MakeBearer(24))
}
