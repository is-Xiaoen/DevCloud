package api

import (
	"122.51.31.227/go-course/go18/devcloud/audit/audit"
	"122.51.31.227/go-course/go18/devcloud/mcenter/permission"
	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func init() {
	ioc.Api().Registry(&UserRestfulApiHandler{})
}

type UserRestfulApiHandler struct {
	ioc.ObjectImpl

	// 依赖控制器
	svc application.Service
}

func (h *UserRestfulApiHandler) Name() string {
	return "applications"
}

func (h *UserRestfulApiHandler) Init() error {
	h.svc = application.GetService()

	tags := []string{"应用管理"}
	ws := gorestful.ObjectRouter(h)

	ws.Route(ws.POST("").To(h.CreateApplication).
		Doc("应用创建").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(permission.Auth(true)).
		Metadata(permission.Permission(true)).
		Metadata(permission.Resource("application")).
		Metadata(permission.Action("create")).
		Metadata(audit.Enable(true)).
		Writes(application.Application{}).
		Returns(200, "OK", application.Application{}))

	// required_auth=true/false
	ws.Route(ws.GET("").To(h.QueryApplication).
		Doc("应用列表查询").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 这个开关怎么生效
		// 中间件需求读取接口的描述信息，来决定是否需要认证
		Metadata(permission.Auth(true)).
		Metadata(permission.Permission(true)).
		Metadata(permission.Resource("application")).
		Metadata(permission.Action("list")).
		Metadata(audit.Enable(true)).
		Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
		Writes(Set{}).
		Returns(200, "OK", Set{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteApplication).
		Doc("应用删除").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(permission.Auth(true)).
		Metadata(permission.Permission(true)).
		Metadata(permission.Resource("application")).
		Metadata(permission.Action("delete")).
		Metadata(audit.Enable(true)).
		Writes().
		Returns(200, "OK", nil))

	return nil
}

// *types.Set[*User]
// 返回的泛型, API Doc这个工具 不支持泛型
type Set struct {
	Total int64                     `json:"total"`
	Items []application.Application `json:"items"`
}
