package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"strings"
	"testing"
)

// 1e9~1e18 √n logn 1     二分 二进制
// 1e5~1e6  nlogn nαn n   RMQ 并查集
// 1e3~1e4  n^2 n√n       RMQ DP 分块
// 300~500  n^3           DP 二分图
func Test_run(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customInputs := []string{
		``,
	}
	customAnswers := []string{
		``,
	}
	if len(customInputs) > 0 && strings.TrimSpace(customInputs[0]) != "" {
		testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
		t.Log("======================================")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCase(t, dir, 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}
