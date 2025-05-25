package generate_test

import (
	"testing"

	"122.51.31.227/go-course/go18/skills/generate"
)

func TestStringSet(t *testing.T) {
	set := generate.NewSet[string]()
	set.Add("test")
	t.Log(set)
}

func TestIntSet(t *testing.T) {
	set := generate.NewSet[int]()
	set.Add(10)
	t.Log(set)
}
