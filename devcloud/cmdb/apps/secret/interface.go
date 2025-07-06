package secret

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/task/apps/task"
)

const (
	APP_NAME   = "secret"
	SECRET_KEY = "23gs6gxHrz1kNEvshRmunkXbwIiaEcYfh+EMu+e9ewA="
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// 用于Secret的管理(后台管理员配置)
	// 创建secret
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	// 查询secret
	QuerySecret(context.Context, *QuerySecretRequest) (*types.Set[*Secret], error)
	// 查询详情, 已解密, API层需要脱敏
	DescribeSecret(context.Context, *DescribeSecretRequeset) (*Secret, error)

	// 基于云商凭证来同步资源
	// 怎么API怎么设计
	// 同步阿里云所有资源, 10分钟，30分钟 ...
	// 这个接口调用持续30分钟...
	// 需要拆解为异步任务: 用户调用了同步后, 里面返回, 这个同步任务在后台执行(Gorouties), 需要查询同步日志(Ws)

	// 使用task模块来执行异步任务, 通过TaskId查询异步任务状态
	SyncResource(context.Context, *SyncResourceRequest) (*types.Set[*task.Task], error)
}

// 同步记录(task) 有状态
type SyncResourceTask struct {
}

type QuerySyncLogRequest struct {
	TaskId int `json:"task_id"`
}

// 每个资源的同步日志
type SyncRecord struct {
}

type ResourceResponse struct {
	Success    bool
	InstanceId string             `json:"instance_id"`
	Resource   *resource.Resource `json:"resource"`
	Message    string             `json:"message"`
}

func (t ResourceResponse) String() string {
	return pretty.ToJSON(t)
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

type QuerySecretRequest struct {
	policy.ResourceScope
	// 分页请求
	*request.PageRequest
}

func NewDescribeSecretRequeset(id string) *DescribeSecretRequeset {
	return &DescribeSecretRequeset{
		Id: id,
	}
}

type DescribeSecretRequeset struct {
	policy.ResourceScope
	Id string `json:"id"`
}

func NewSyncResourceRequest(id string) *SyncResourceRequest {
	return &SyncResourceRequest{
		Id: id,
	}
}

type SyncResourceRequest struct {
	Id string `json:"id"`
}
