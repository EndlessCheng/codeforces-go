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
			``,
			``,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
