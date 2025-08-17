package api

import (
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/label"
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin/binding"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func (h *LabelRestfulApiHandler) QueryLabel(r *restful.Request, w *restful.Response) {
	req := label.NewQueryLabelRequest()

	// 过滤条件在认证完成后的上下文中
	tk := token.GetTokenFromCtx(r.Request.Context())
	log.L().Debug().Msgf("resource scope: %s", tk.ResourceScope)

	// 用户的参数
	if err := binding.Query.Bind(r.Request, &req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.svc.QueryLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
