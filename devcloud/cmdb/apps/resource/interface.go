package resource

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

const (
	APP_NAME = "resource"
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// 资源搜索
	Search(context.Context, *SearchRequest) (*types.Set[*Resource], error)
	// Save 更新与创建同时
	Add(context.Context, *Resource) (*Resource, error)
	// 删除
	DeleteResource(context.Context, *DeleteResourceRequest) error
}

// NewSearchRequest().SetType(”).SetXXX(v)
// (WithOptz(), WithOptx(), WithOpty()...)
func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		PageRequest: request.NewDefaultPageRequest(),
		Tags:        map[string]string{},
	}
}

type SearchRequest struct {
	// 分页请求
	*request.PageRequest
	// 名称做模糊搜索
	Keywords string `json:"keywords"`
	// 类型
	Type *TYPE `json:"type"`
	// 标签
	Tags map[string]string `json:"lable"`
}

func (r *SearchRequest) SetType(t TYPE) *SearchRequest {
	r.Type = &t
	return r
}

func NewDeleteResourceRequest() *DeleteResourceRequest {
	return &DeleteResourceRequest{}
}

// 支持多个
type DeleteResourceRequest struct {
	ResourceIds []string `json:"resource_ids"`
}
