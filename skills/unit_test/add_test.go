package unittest_test

import (
	"os"
	"testing"

	unittest "122.51.31.227/go-course/go18/skills/unit_test"
	"github.com/stretchr/testify/assert"
)

// 针对Add函数的单元测试
func TestAdd(t *testing.T) {
	t.Log(os.Getenv("CONFIG_PATH"))
	t.Log(os.Getenv("DB_HOST"))
	// 自己手动判断
	// /usr/local/go/bin/go test -timeout 300s -run ^TestAdd$ 122.51.31.227/go-course/go18/skills/unit_test -v -count=1
	// 如果没有打印日志， 配置vscode 打印单元测试的日志: -v -count=1
	// 加上这串配置
	//     "go.testFlags": [
	//     "-v",
	//     "-count=1"
	// ],
	t.Log(unittest.Add(1, 2))

	// 通过程序断言来判断
	// if unittest.Add(1, 2) != 4 {
	// 	t.Fatalf("1 + 2 != 4")
	// }
	// 专门的断言库
	should := assert.New(t)
	should.Equal(3, unittest.Add(1, 2))
}
