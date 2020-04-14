package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customTestCases := [][2]string{
		{
			``,
			``,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		testutil.AssertEqualStringCase(t, customTestCases, 0, run)
		//testutil.AssertEqualRunResults(t, customTestCases, 0, runAC, run)
		t.Log("======= custom =======")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCase(t, dir, 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}
