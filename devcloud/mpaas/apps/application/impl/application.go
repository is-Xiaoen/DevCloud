package impl

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"gorm.io/gorm"
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

	i.log.Debug().Msgf("scope: %s", in.ResourceScope)
	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&application.Application{}))
	if in.Id != "" {
		query = query.Where("id = ?", in.Id)
	}
	if in.Name != "" {
		query = query.Where("name = ?", in.Name)
	}
	if in.Ready != nil {
		query = query.Where("ready = ?", *in.Ready)
	}
	if in.Keywords != "" {
		query = query.Where("name LIKE ?", "%"+in.Keywords+"%")
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
func (i *ApplicationServiceImpl) DescribeApplication(ctx context.Context, in *application.DescribeApplicationRequest) (*application.Application, error) {
	query := in.GormResourceFilter(datasource.DBFromCtx(ctx).Model(&application.Application{}))

	ins := &application.Application{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("app %s not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}

// UpdateApplication implements application.Service.
func (i *ApplicationServiceImpl) UpdateApplication(ctx context.Context, in *application.UpdateApplicationRequest) (*application.Application, error) {
	ins, err := i.DescribeApplication(ctx, &in.DescribeApplicationRequest)
	if err != nil {
		return nil, err
	}

	ins.CreateApplicationSpec = in.CreateApplicationSpec
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

// DeleteApplication implements application.Service.
func (i *ApplicationServiceImpl) DeleteApplication(ctx context.Context, in *application.DeleteApplicationRequest) (*application.Application, error) {
	ins, err := i.DescribeApplication(ctx, &in.DescribeApplicationRequest)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).
		Where("id = ?", in.Id).
		Delete(&application.Application{}).
		Error
}
