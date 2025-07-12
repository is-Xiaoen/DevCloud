package impl

import (
	"context"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
)

// Add implements resource.Service.
func (s *ResourceServiceImpl) Add(ctx context.Context, in *resource.Resource) (*resource.Resource, error) {
	if err := datasource.DBFromCtx(ctx).Save(in).Error; err != nil {
		return nil, err
	}

	return in, nil
}

// DeleteResource implements resource.Service.
func (s *ResourceServiceImpl) DeleteResource(ctx context.Context, in *resource.DeleteResourceRequest) error {
	panic("unimplemented")
}

// Search implements resource.Service.
func (s *ResourceServiceImpl) Search(ctx context.Context, in *resource.SearchRequest) (*types.Set[*resource.Resource], error) {
	panic("unimplemented")
}
