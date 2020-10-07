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
			`3 3 1 4
3 5
1 2
2 6
1 6
3 5
1 4`,
			`10
2
1
-1`,
		},
		{
			`1 2 1 5
1 5
1 8
1 4`,
			`10
-1
1`,
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
// https://www.luogu.com.cn/problem/P6851
