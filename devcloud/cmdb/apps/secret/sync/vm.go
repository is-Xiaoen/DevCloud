package sync

import (
	"context"
	"fmt"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/mcube/v2/tools/ptr"
	"github.com/infraboard/modules/task/apps/event"
	"github.com/infraboard/modules/task/apps/task"
	"github.com/rs/zerolog"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

func init() {
	ioc.Controller().Registry(&VmSyncerServiceImpl{})
}

var _ secret.Syncer = (*VmSyncerServiceImpl)(nil)

type VmSyncerServiceImpl struct {
	ioc.ObjectImpl

	log *zerolog.Logger
}

func (s *VmSyncerServiceImpl) Name() string {
	return secret.GetSyncerName(resource.TYPE_VM)
}

func (s *VmSyncerServiceImpl) Init() error {
	s.log = log.Sub(s.Name())
	secret.RegistrySyncer(resource.TYPE_VM, s)
	return nil
}

// Sync implements secret.Syncer.
func (s *VmSyncerServiceImpl) Sync(ctx context.Context, ins *secret.Secret, region string, rs resource.TYPE) *task.Task {
	// 怎么使用对应secret 来完成, 用的腾讯的API
	// https://console.cloud.tencent.com/api/explorer?Product=cvm&Version=2017-03-12&Action=DescribeRegions
	// https://console.cloud.tencent.com/api/explorer?Product=lighthouse&Version=2020-03-24&Action=DescribeInstances
	return task.GetService().Run(ctx, task.NewFnTask(func(ctx context.Context, req any) error {
		taskInfo := task.GetTaskFromCtx(ctx)
		_, err := event.GetService().AddEvent(ctx, task.NewInfoEvent("同步开始", taskInfo.Id))
		if err != nil {
			return err
		}
		defer event.GetService().AddEvent(ctx, task.NewInfoEvent("同步结束", taskInfo.Id))

		// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
		// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性
		// 以下代码示例仅供参考，建议采用更安全的方式来使用密钥
		// 请参见：https://cloud.tencent.com/document/product/1278/85305
		// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
		credential := common.NewCredential(
			ins.ApiKey,
			ins.ApiSecret,
		)
		// 使用临时密钥示例
		// credential := common.NewTokenCredential("SecretId", "SecretKey", "Token")
		// 实例化一个client选项，可选的，没有特殊需求可以跳过
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
		// 实例化要请求产品的client对象,clientProfile是可选的
		// "ap-shanghai"
		client, _ := lighthouse.NewClient(credential, region, cpf)

		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := lighthouse.NewDescribeInstancesRequest()

		// 返回的resp是一个DescribeInstancesResponse的实例，与请求对象对应
		response, err := client.DescribeInstances(request)
		if err != nil {
			return err
		}

		_, err = event.GetService().AddEvent(ctx, task.NewInfoEvent("接口调用成功", taskInfo.Id).SetDetail(response.ToJsonString()))
		if err != nil {
			return err
		}

		// 获取出了实例的列表
		for _, vm := range response.Response.InstanceSet {
			// 转化为一个Resource对象
			res := resource.NewResource()
			res.Vendor = resource.VENDOR_TENCENT
			res.CredentialId = ins.Id
			res.Region = region
			res.ResourceType = resource.TYPE_VM
			res.Id = ptr.GetValue(vm.InstanceId)
			res.Name = ptr.GetValue(vm.InstanceName)
			res.Cpu = ptr.GetValue(vm.CPU)
			res.Memory = ptr.GetValue(vm.Memory)
			res.SystemStorage = ptr.GetValue(vm.SystemDisk.DiskSize)
			res.Extra["os_name"] = ptr.GetValue(vm.OsName)
			res.PrivateAddress = ptr.GetArrayValue(vm.PrivateAddresses)
			res.PublicAddress = ptr.GetArrayValue(vm.PublicAddresses)
			res.Phase = ptr.GetValue(vm.InstanceState)
			res, err := resource.GetService().Add(ctx, res)
			if err != nil {
				return err
			}
			_, err = event.GetService().AddEvent(ctx, task.NewInfoEvent(fmt.Sprintf("资源%s同步成功", res.Name), taskInfo.Id).SetDetail(res.String()))
			if err != nil {
				return err
			}
		}
		return nil
	}, ins).SetAsync(true).SetLabel(secret.TASK_LABLE_SECRET_ID, ins.Id))
}
