package config_test

import (
	"fmt"
	"os"
	"testing"

	"122.51.31.227/go-course/go18/book/v3/config"
)

func TestLoadConfigFromYaml(t *testing.T) {
	err := config.LoadConfigFromYaml(fmt.Sprintf("%s/book/v2/application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_HOST", "localhost")
	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}
