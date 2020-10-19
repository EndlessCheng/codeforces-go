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
			`3
AKICKSTARTPROBLEMNAMEDKICKSTART
STARTUNLUCKYKICK
KICKXKICKXSTARTXKICKXSTART`,
			`Case #1: 3
Case #2: 0
Case #3: 5`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
