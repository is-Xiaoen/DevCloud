package impl

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
)

// CreateApplication implements application.Service.
func (i *ApplicationServiceImpl) CreateApplication(ctx context.Context, in *application.CreateApplicationRequest) (*application.Application, error) {
	ins, err := application.NewApplication(*in)
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).
		Save(ins).
		Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// QueryApplication implements application.Service.
func (i *ApplicationServiceImpl) QueryApplication(ctx context.Context, in *application.QueryApplicationRequest) (*types.Set[*application.Application], error) {
	set := types.New[*application.Application]()

	query := datasource.DBFromCtx(ctx).Model(&application.Application{})
	if in.Id != "" {
		query = query.Where("id = ?", in.Id)
	}
	if in.Name != "" {
		query = query.Where("name = ?", in.Name)
	}
	if in.Ready != nil {
		query = query.Where("ready = ?", *in.Ready)
	}

	if in.NamespaceId != nil {
		query = query.Where("namespace = ?", in.NamespaceId)
	}
	// 过滤条件, Label

	if in.Scope != nil {
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("created_at desc").
		Offset(int(in.ComputeOffset())).
		Limit(int(in.PageSize)).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// DescribeApplication implements application.Service.
func (i *ApplicationServiceImpl) DescribeApplication(context.Context, *application.DescribeApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}

// UpdateApplication implements application.Service.
func (i *ApplicationServiceImpl) UpdateApplication(context.Context, *application.UpdateApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}

// DeleteApplication implements application.Service.
func (i *ApplicationServiceImpl) DeleteApplication(context.Context, *application.DeleteApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}
