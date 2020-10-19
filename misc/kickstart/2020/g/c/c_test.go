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
3 5
2 3 4
4 10
2 9 3 8`,
			`Case #1: 2
Case #2: 8`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
