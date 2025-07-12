package secret

import (
	"context"
	"fmt"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"github.com/infraboard/modules/task/apps/task"
)

const (
	SYNCER_PREFIX = "syncer"
)

var (
	syncers = map[string]Syncer{}
)

func GetSyncerName(t resource.TYPE) string {
	return fmt.Sprintf("%s_%s", SYNCER_PREFIX, t)
}

func GetSyncer(t resource.TYPE) Syncer {
	return syncers[fmt.Sprintf("%s_%s", SYNCER_PREFIX, t)]
}

func RegistrySyncer(t resource.TYPE, s Syncer) {
	syncers[fmt.Sprintf("%s_%s", SYNCER_PREFIX, t)] = s
}

type Syncer interface {
	// 资源同步
	Sync(ctx context.Context, ins *Secret, region string, rt resource.TYPE) *task.Task
}
