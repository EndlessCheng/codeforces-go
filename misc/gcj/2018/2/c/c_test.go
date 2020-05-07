package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [C]")
	testCases := [][2]string{
		{
			`4
2
1 2
2 1
2
1 1
2 1
2
1 2
1 2
2
2 2
-2 2`,
			`Case #1: 0
Case #2: 1
Case #3: 2
Case #4: 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
