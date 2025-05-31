package token

import (
	"context"
	"time"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

const (
	APP_NAME = "token"
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// 颁发访问令牌: Login
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销访问令牌: 令牌失效了 Logout
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	// 查询已经颁发出去的Token
	QueryToken(context.Context, *QueryTokenRequest) (*types.Set[*Token], error)

	// 查询Token详情
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
	// 校验访问令牌：检查令牌的合法性, 是不是伪造的
	ValiateToken(context.Context, *ValiateTokenRequest) (*Token, error)
}

func NewDescribeTokenRequest(accessToken string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		DescribeBy:    DESCRIBE_BY_ACCESS_TOKEN,
		DescribeValue: accessToken,
	}
}

type DescribeTokenRequest struct {
	DescribeBy    DESCRIBE_BY `json:"describe_by"`
	DescribeValue string      `json:"describe_value"`
}

func NewQueryTokenRequest() *QueryTokenRequest {
	return &QueryTokenRequest{
		PageRequest: request.NewDefaultPageRequest(),
		UserIds:     []uint64{},
	}
}

type QueryTokenRequest struct {
	*request.PageRequest
	// 当前可用的没过期的Token
	Active *bool `json:"active"`
	// 用户来源
	Source *SOURCE `json:"source"`
	// Uids
	UserIds []uint64 `json:"user_ids"`
}

func (r *QueryTokenRequest) SetActive(v bool) *QueryTokenRequest {
	r.Active = &v
	return r
}

func (r *QueryTokenRequest) SetSource(v SOURCE) *QueryTokenRequest {
	r.Source = &v
	return r
}

func (r *QueryTokenRequest) AddUserId(uids ...uint64) *QueryTokenRequest {
	r.UserIds = append(r.UserIds, uids...)
	return r
}

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		Parameter: make(IssueParameter),
	}
}

// 用户会给我们 用户的身份凭证，用于换取Token
type IssueTokenRequest struct {
	// 端类型
	Source SOURCE `json:"source"`
	// 认证方式
	Issuer string `json:"issuer"`
	// 参数
	Parameter IssueParameter `json:"parameter"`
}

func (i *IssueTokenRequest) IssueByPassword(username, password string) {
	i.Issuer = ISSUER_PASSWORD
	i.Parameter.SetUsername(username)
	i.Parameter.SetPassword(password)
}

func NewIssueParameter() IssueParameter {
	return make(IssueParameter)
}

type IssueParameter map[string]any

/*
password issuer parameter
*/

func (p IssueParameter) Username() string {
	return GetIssueParameterValue[string](p, "username")
}

func (p IssueParameter) Password() string {
	return GetIssueParameterValue[string](p, "password")
}

func (p IssueParameter) SetUsername(v string) IssueParameter {
	p["username"] = v
	return p
}

func (p IssueParameter) SetPassword(v string) IssueParameter {
	p["password"] = v
	return p
}

/*
private token issuer parameter
*/

func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

func (p IssueParameter) SetAccessToken(v string) IssueParameter {
	p["access_token"] = v
	return p
}

func (p IssueParameter) SetExpireTTL(v int64) IssueParameter {
	p["expired_ttl"] = v
	return p
}

func NewRevolkTokenRequest(at, rk string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  at,
		RefreshToken: rk,
	}
}

// 万一的Token泄露, 不知道refresh_token，也没法推出
type RevolkTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewValiateTokenRequest(accessToken string) *ValiateTokenRequest {
	return &ValiateTokenRequest{
		AccessToken: accessToken,
	}
}

type ValiateTokenRequest struct {
	AccessToken string `json:"access_token"`
}
