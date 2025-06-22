package impl_test

import (
	"testing"

	"122.51.31.227/go-course/go18/devcloud/mpaas/apps/application"
)

func TestCreateApplication(t *testing.T) {
	req := application.NewCreateApplicationRequest()
	req.Name = "devcloud"
	req.Description = "应用研发云"
	req.Type = application.TYPE_SOURCE_CODE
	req.CodeRepository = application.CodeRepository{
		SshUrl: "git@122.51.31.227:go-course/go18.git",
	}
	req.SetLabel("team", "golang_dev_group_01")
	ins, err := svc.CreateApplication(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// SELECT * FROM `applications` WHERE (NOT JSON_EXTRACT(`label`,'$.team') IS NOT NULL OR JSON_EXTRACT(`label`,'$.team') = 'golang_dev_group_01') ORDER BY created_at desc LIMIT 20
func TestQueryApplication(t *testing.T) {
	req := application.NewQueryApplicationRequest()
	req.SetScope("team", []string{"golang_dev_group_01"})
	// req.SetScope("env", []string{"prod"})
	ins, err := svc.QueryApplication(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
