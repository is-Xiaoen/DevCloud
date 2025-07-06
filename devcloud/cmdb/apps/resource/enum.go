package resource

const (
	VENDOR_ALIYUN  VENDOR = "ali"
	VENDOR_TENCENT VENDOR = "tx"
	VENDOR_HUAWEI  VENDOR = "hw"
	VENDOR_VMWARE  VENDOR = "vmware"
	VENDOR_AWS     VENDOR = "aws"
)

type VENDOR string

const (
	// 业务资源
	TYPE_VM      TYPE = "vm"
	TYPE_RDS     TYPE = "rds"
	TYPE_REDIS   TYPE = "redis"
	TYPE_BUCKET  TYPE = "bucket"
	TYPE_DISK    TYPE = "disk"
	TYPE_LB      TYPE = "lb"
	TYPE_DOMAIN  TYPE = "domain"
	TYPE_EIP     TYPE = "eip"
	TYPE_MONGODB TYPE = "mongodb"
	// 子资源
	TYPE_DATABASE TYPE = "database"
	TYPE_ACCOUNT  TYPE = "account"
	// 未知资源
	TYPE_OTHER TYPE = "other"
	// 辅助资源
	TYPE_BILL  TYPE = "bill"
	TYPE_ORDER TYPE = "order"
)

type TYPE string

const ()

type STATUS string
