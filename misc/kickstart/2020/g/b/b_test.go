package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	dir, _ := filepath.Abs(".")
	t.Logf("Current problem is [%s]", filepath.Base(dir))

	customTestCases := [][2]string{
		{
			`2
3
1 2 5
3 6 1
12 2 7
5
0 0 0 0 0
1 1 1 1 0
2 2 2 8 0
1 1 1 0 0
0 0 0 0 0`,
			`Case #1: 14
Case #2: 9`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
