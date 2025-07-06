package impl_test

import (
	"testing"
	"time"

	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/resource"
	"122.51.31.227/go-course/go18/devcloud/cmdb/apps/secret"
)

func TestCreateSecret(t *testing.T) {
	req := secret.NewCreateSecretRequest()
	req.Name = "腾讯云只读账号"
	req.Vendor = resource.VENDOR_TENCENT
	req.ApiKey = "xxx"
	req.ApiSecret = "xx"
	req.SetEnabled(true)
	req.ResourceType = append(req.ResourceType, resource.TYPE_VM)
	req.Regions = []string{"ap-shanghai", "ap-guangzhou"}
	ins, err := svc.CreateSecret(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQuerySecret(t *testing.T) {
	req := secret.NewQuerySecretRequest()
	set, err := svc.QuerySecret(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

const (
	SECRET_ID = "ed02f7cd-ee5b-33f0-bdf5-e305ecb3efb4"
)

func TestDescribeSecret(t *testing.T) {
	req := secret.NewDescribeSecretRequeset(SECRET_ID)
	ins, err := svc.DescribeSecret(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestSyncResource(t *testing.T) {
	req := secret.NewSyncResourceRequest(SECRET_ID)
	ins, err := svc.SyncResource(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

	time.Sleep(3 * time.Second)
}
