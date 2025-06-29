package impl_test

import (
	"testing"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/policy"
)

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.WithUser = true
	req.WithRole = true
	req.WithNamespace = true
	set, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	// guest
	req.UserId = 2
	// 开发
	req.RoleId = []uint64{3}
	// default
	req.SetNamespaceId(1)
	// 开发小组1的资源
	req.SetScope("team", []string{"dev01"})
	set, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
