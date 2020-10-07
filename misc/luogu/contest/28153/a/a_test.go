package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	customTestCases := [][2]string{
		{
			`50 50 72 56 100 40 0 1 446`,
			`AFO`,
		},
		{
			`50 95 100 64 100 72 30 0 446`,
			`AKIOI`,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		testutil.AssertEqualStringCase(t, customTestCases, 0, run)
		//testutil.AssertEqualRunResults(t, customTestCases, 0, runAC, run)
		t.Log("======= custom =======")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, run)
	//testutil.AssertEqualFileCaseWithName(t, dir, "*.in", "*.out", 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}
// https://www.luogu.com.cn/problem/P6850
