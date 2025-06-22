package impl_test

import (
	"testing"

	"122.51.31.227/go-course/go18/devcloud/mcenter/apps/label"
)

func TestCreateLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Key = "tech"
	req.KeyDesc = "技术部"
	req.ValueType = label.VALUE_TYPE_ENUM
	req.AddEnumOption(&label.EnumOption{
		Label: "开发一组",
		Value: "dev01",
	})

	ins, err := svc.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
