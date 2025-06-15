package impl

import (
	"122.51.31.227/go-course/go18/devcloud/audit/apps/event"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"

	// 引入消费者
	_ "122.51.31.227/go-course/go18/devcloud/audit/apps/event/consumer"
)

func init() {
	ioc.Controller().Registry(&EventServiceImpl{})
}

var _ event.Service = (*EventServiceImpl)(nil)

// 业务具体实现
type EventServiceImpl struct {
	// 继承模版
	ioc.ObjectImpl

	// 模块子Logger
	log *zerolog.Logger

	// MongoDB集合
	col *mongo.Collection
}

// 对象名称
func (i *EventServiceImpl) Name() string {
	return event.AppName
}

func (i *EventServiceImpl) Priority() int {
	return event.PRIORITY
}

// 初始化
func (i *EventServiceImpl) Init() error {
	// 对象
	i.log = log.Sub(i.Name())

	i.log.Debug().Msgf("database: %s", ioc_mongo.Get().Database)
	// 需要一个集合Collection
	i.col = ioc_mongo.DB().Collection("events")
	return nil
}
