package token

import "github.com/infraboard/mcube/v2/exception"

const (
	ACCESS_TOKEN_HEADER_NAME          = "Authorization"
	ACCESS_TOKEN_COOKIE_NAME          = "access_token"
	ACCESS_TOKEN_RESPONSE_HEADER_NAME = "X-OAUTH-TOKEN"
	REFRESH_TOKEN_HEADER_NAME         = "X-REFRUSH-TOKEN"
)

// 自定义非导出类型，避免外部包直接实例化
type tokenContextKey struct{}

var (
	CTX_TOKEN_KEY = tokenContextKey{}
)

var (
	CookieNotFound = exception.NewUnauthorized("cookie %s not found", ACCESS_TOKEN_COOKIE_NAME)
)
