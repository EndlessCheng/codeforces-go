package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`4
1 3
1
5 2
10 5 359999999999 123456789 10
2 3
8 4
3 2
1 2 3`,
			`Case #1: 2
Case #2: 0
Case #3: 1
Case #4: 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
