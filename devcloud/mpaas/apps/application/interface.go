package application

import (
	"context"

	"github.com/infraboard/mcube/v2/types"
)

type Service interface {
	// 创建应用
	CreateApplication(context.Context, CreateApplicationRequest) (*Application, error)
	// 查询应用
	QueryApplication(context.Context, QueryApplicationRequest) (*types.Set[*Application], error)
	// 更新应用
	UpdateApplication(context.Context, UpdateApplicationRequest) (*Application, error)
	// 删除应用
	DeleteApplication(context.Context, DeleteApplicationRequest) (*Application, error)
	// 获取应用
	DescribeApplication(context.Context, DescribeApplicationRequest) (*Application, error)
}

type QueryApplicationRequest struct {
	// 应用ID
	Id string `json:"id" bson:"_id"`
	// 应用名称
	Name string `json:"name" bson:"name"`
	// 应用状态
	Status string `json:"status" bson:"status"`
}

type UpdateApplicationRequest struct {
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by"`
	DescribeApplicationRequest
	CreateApplicationSpec
}

type DeleteApplicationRequest struct {
	DescribeApplicationRequest
}

type DescribeApplicationRequest struct {
	// 应用ID
	Id string `json:"id" bson:"_id"`
}
