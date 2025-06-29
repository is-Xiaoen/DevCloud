package api

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/token"
	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin/binding"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func (h *UserRestfulApiHandler) QueryApplication(r *restful.Request, w *restful.Response) {
	req := application.NewQueryApplicationRequest()
	if err := binding.Query.Bind(r.Request, &req.QueryApplicationRequestSpec); err != nil {
		response.Failed(w, err)
		return
	}

	// 过滤条件在认证完成后的上下文中
	tk := token.GetTokenFromCtx(r.Request.Context())
	req.ResourceScope = tk.ResourceScope
	log.L().Debug().Msgf("resource scope: %s", tk.ResourceScope)

	set, err := h.svc.QueryApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
