package application

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

const (
	APP_NAME = "application"
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// 创建应用
	CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error)
	// 查询应用
	QueryApplication(context.Context, *QueryApplicationRequest) (*types.Set[*Application], error)
	// 更新应用
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)
	// 删除应用
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*Application, error)
	// 获取应用
	DescribeApplication(context.Context, *DescribeApplicationRequest) (*Application, error)
}

func NewQueryApplicationRequest() *QueryApplicationRequest {
	return &QueryApplicationRequest{
		QueryApplicationRequestSpec: QueryApplicationRequestSpec{
			PageRequest: request.NewDefaultPageRequest(),
		},
	}
}

type QueryApplicationRequest struct {
	policy.ResourceScope
	QueryApplicationRequestSpec
}

type QueryApplicationRequestSpec struct {
	*request.PageRequest
	// 应用ID
	Id string `json:"id" form:"id"`
	// 应用名称
	Name string `json:"name" form:"name"`
	// 应用是否就绪
	Ready *bool `json:"ready" form:"ready"`
	// 关键字
	Keywords string `json:"keywords" form:"keywords"`
}

type UpdateApplicationRequest struct {
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by"`
	DescribeApplicationRequest
	CreateApplicationSpec
}

func NewDeleteApplicationRequest(appId string) *DeleteApplicationRequest {
	return &DeleteApplicationRequest{
		DescribeApplicationRequest: *NewDescribeApplicationRequest(appId),
	}
}

type DeleteApplicationRequest struct {
	DescribeApplicationRequest
}

func NewDescribeApplicationRequest(appId string) *DescribeApplicationRequest {
	return &DescribeApplicationRequest{
		Id: appId,
	}
}

type DescribeApplicationRequest struct {
	policy.ResourceScope
	// 应用ID
	Id string `json:"id" bson:"_id"`
}
