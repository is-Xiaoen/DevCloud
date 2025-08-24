package hello_service_test

import (
	"testing"

	"122.51.31.227/go-course/go18/skills/rpc/protobuf/hello_service"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestAny(t *testing.T) {
	status := hello_service.ErrorStatus{
		Details: []*anypb.Any{},
	}

	d1 := hello_service.Request{Value: "test"}
	any1, err := anypb.New(&d1)
	if err != nil {
		t.Fatal(err)
	}

	d2 := hello_service.Response{Value: "test"}
	any2, err := anypb.New(&d2)
	if err != nil {
		t.Fatal(err)
	}
	status.Details = append(status.Details, any1, any2)

	t.Log(status)

	target := hello_service.Request{}
	if err := status.Details[0].UnmarshalTo(&target); err != nil {
		t.Fatal(err)
	}
	t.Log(target.Value)
}
