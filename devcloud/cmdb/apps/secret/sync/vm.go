package sync

import (
	"context"
	"time"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/modules/task/apps/task"
	"github.com/rs/zerolog"
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
func (s *VmSyncerServiceImpl) Sync(ctx context.Context, secretIns *secret.Secret, rs resource.TYPE) *task.Task {
	return task.GetService().Run(ctx, task.NewFnTask(func(ctx context.Context, req any) error {

		// taskInfo := task.GetTaskFromCtx(ctx)
		s.log.Debug().Msg("test for vm sync wait")
		time.Sleep(1 * time.Second)
		// secrt同步
		return nil
	}, secretIns).SetAsync(true).SetLabel(secret.TASK_LABLE_SECRET_ID, secretIns.Id))
}
