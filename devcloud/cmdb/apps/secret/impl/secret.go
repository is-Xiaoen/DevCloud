package impl

import (
	"context"
	"fmt"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/task/apps/task"
	"gorm.io/gorm"
)

// CreateSecret implements secret.Service.
func (s *SecretServiceImpl) CreateSecret(ctx context.Context, in *secret.CreateSecretRequest) (*secret.Secret, error) {
	ins := secret.NewSecret(in)

	// 需要加密
	if err := ins.EncryptedApiSecret(); err != nil {
		return nil, err
	}

	// upsert, gorm save
	if err := datasource.DBFromCtx(ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

// DescribeSecret implements secret.Service.
func (s *SecretServiceImpl) DescribeSecret(ctx context.Context, in *secret.DescribeSecretRequeset) (*secret.Secret, error) {
	// 取出后，需要解密
	ins := &secret.Secret{}
	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&secret.Secret{}))
	if err := query.Where("id = ?", in.Id).Take(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("secret %s not found", in.Id)
		}
		return nil, err
	}
	ins.SetIsEncrypted(true)
	if err := ins.DecryptedApiSecret(); err != nil {
		return nil, err
	}

	// 解密过后的数据
	return ins, nil
}

// QuerySecret implements secret.Service.
func (s *SecretServiceImpl) QuerySecret(ctx context.Context, in *secret.QuerySecretRequest) (*types.Set[*secret.Secret], error) {
	set := types.New[*secret.Secret]()

	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&secret.Secret{}))
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("create_at desc").
		Offset(int(in.ComputeOffset())).
		Limit(int(in.PageSize)).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}
	return set, nil

}

// SyncResource implements secret.Service.
// 资源同步接口
func (s *SecretServiceImpl) SyncResource(ctx context.Context, in *secret.SyncResourceRequest) (*types.Set[*task.Task], error) {
	// 直接初始化并返回，避免警告
	taskSet := types.NewSet[*task.Task]()

	// 查询Secret信息
	ins, err := s.DescribeSecret(ctx, secret.NewDescribeSecretRequeset(in.Id))
	if err != nil {
		return taskSet, err
	}

	if !ins.GetEnabled() {
		return taskSet, fmt.Errorf("secret %s not enabled", ins.Name)
	}

	// 获取syncer,  执行同步
	for _, rs := range ins.ResourceType {
		syncer := secret.GetSyncer(rs)
		for _, region := range ins.Regions {
			taskInfo := syncer.Sync(ctx, ins, region, rs)
			taskSet.Add(taskInfo)
		}
	}
	return taskSet, nil
}
