package resource

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
)

func NewResourceSet() *types.Set[*Resource] {
	return types.New[*Resource]()
}

func NewResource() *Resource {
	return &Resource{
		Meta: Meta{},
		Spec: Spec{
			Tags:  map[string]string{},
			Extra: map[string]string{},
		},
		Status: Status{},
	}
}

// 资源
// https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/#struct-tags
type Resource struct {
	Meta   `bson:"inline"`
	Spec   `bson:"inline"`
	Status `bson:"inline"`
}

func (r *Resource) String() string {
	return pretty.ToJSON(r)
}

func (r *Resource) TableName() string {
	return "resources"
}

// 元数据，不会变的
type Meta struct {
	// 全局唯一Id, 直接使用个云商自己的Id
	Id string `bson:"_id" json:"id" validate:"required" gorm:"column:id;primaryKey"`
	// 资源所属域
	Domain string `json:"domain" validate:"required" gorm:"column:domain"`
	// 资源所属空间
	Namespace string `json:"namespace" validate:"required" gorm:"column:namespace"`
	// 资源所属环境
	Env string `json:"env" gorm:"column:env"`
	// 创建时间
	CreateAt int64 `json:"create_at" gorm:"column:create_at"`
	// 删除时间
	DeteteAt int64 `json:"detete_at" gorm:"column:detete_at"`
	// 删除人
	DeteteBy string `json:"detete_by" gorm:"column:detete_by"`
	// 同步时间
	SyncAt int64 `json:"sync_at" validate:"required" gorm:"column:sync_at"`
	// 同步人
	SyncBy string `json:"sync_by" gorm:"column:sync_by"`
	// 用于同步的凭证ID
	CredentialId string `json:"credential_id" gorm:"column:credential_id"`
	// 序列号
	SerialNumber string `json:"serial_number" gorm:"column:serial_number"`
}

// 表单数据, 用户申请资源时 需要提供的参数
type Spec struct {
	// 厂商
	Vendor VENDOR `json:"vendor" gorm:"column:vendor"`
	// 资源类型
	ResourceType TYPE `json:"resource_type" gorm:"column:resource_type"`
	// 地域
	Region string `json:"region" gorm:"column:region"`
	// 区域
	Zone string `json:"zone" gorm:"column:zone"`
	// 资源所属账号
	Owner string `json:"owner" gorm:"column:owner"`
	// 名称
	Name string `json:"name" gorm:"column:name"`
	// 种类
	Category string `json:"category" gorm:"column:category"`
	// 规格
	Type string `json:"type" gorm:"column:type"`
	// 描述
	Description string `json:"description" gorm:"column:description"`
	// 过期时间
	ExpireAt int64 `json:"expire_at" gorm:"column:expire_at"`
	// 更新时间
	UpdateAt int64 `json:"update_at" gorm:"column:update_at"`
	// 资源占用Cpu数量
	Cpu int64 `json:"cpu" gorm:"column:cpu"`
	// GPU数量
	Gpu int64 `json:"gpu" gorm:"column:gpu"`
	// 资源使用的内存
	Memory int64 `json:"memory" gorm:"column:memory"`
	// 系统盘
	SystemStorage int64 `json:"system_storage" gorm:"column:system_storage"`
	// 数据盘
	DataStorage int64 `json:"data_storage" gorm:"column:data_storage"`
	// 公网IP带宽, 单位M
	BandWidth int32 `json:"band_width" gorm:"column:band_width"`
	// 资源标签
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json;type:json"`
	// 额外的通用属性
	Extra map[string]string `json:"extra" gorm:"column:extra;serializer:json;type:json"`
}

// 资源当前状态
type Status struct {
	// 资源当前状态
	Phase string `json:"phase" gorm:"column:phase"`
	// 资源当前状态描述
	Describe string `json:"describe" gorm:"column:describe"`
	// 实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定
	LockMode string `json:"lock_mode" gorm:"column:lock_mode"`
	// 锁定原因
	LockReason string `json:"lock_reason" gorm:"column:lock_reason"`
	// 资源访问地址
	// 公网地址, 或者域名
	PublicAddress []string `json:"public_address" gorm:"column:public_address;serializer:json;type:json"`
	// 内网地址, 或者域名
	PrivateAddress []string `json:"private_address" gorm:"column:private_address;serializer:json;type:json"`
}

func (s *Status) GetFirstPrivateAddress() string {
	if len(s.PrivateAddress) > 0 {
		return s.PrivateAddress[0]
	}

	return ""
}
