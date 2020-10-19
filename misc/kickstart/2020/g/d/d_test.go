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
2 1 10
5
19 3 78 2 31`,
			`Case #1: 20.00000000
Case #2: 352.33333333`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
