package api

import (
	"122.51.31.227/go-course/go18/devcloud/audit/audit"
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/label"
	"122.51.31.227/go-course/go18/devcloud/mcenter/permission"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func init() {
	ioc.Api().Registry(&LabelRestfulApiHandler{})
}

type LabelRestfulApiHandler struct {
	ioc.ObjectImpl

	// 依赖控制器
	svc label.Service
}

func (h *LabelRestfulApiHandler) Name() string {
	return "labels"
}

func (h *LabelRestfulApiHandler) Init() error {
	h.svc = label.GetService()

	tags := []string{"标签管理"}
	ws := gorestful.ObjectRouter(h)
	// required_auth=true/false
	ws.Route(ws.GET("").To(h.QueryLabel).
		Doc("标签列表查询").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 这个开关怎么生效
		// 中间件需求读取接口的描述信息，来决定是否需要认证
		Metadata(permission.Auth(true)).
		Metadata(permission.Permission(false)).
		Metadata(permission.Resource("labels")).
		Metadata(permission.Action("list")).
		Metadata(audit.Enable(true)).
		Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
		Writes(Set{}).
		Returns(200, "OK", Set{}))
	return nil
}

// *types.Set[*Label]
// 返回的泛型, API Doc这个工具 不支持泛型
type Set struct {
	Total int64         `json:"total"`
	Items []label.Label `json:"items"`
}
