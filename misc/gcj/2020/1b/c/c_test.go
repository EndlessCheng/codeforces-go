package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`
1
5 3`,
			``,
		},
		{
			`
9
4 5
4 4
4 3
4 2
5 6
5 5
5 4
5 3
5 2`,
			``,
		},
		{
			`
3
2 2
3 2
2 3`,
			`
Case #1: 1
2 1
Case #2: 2
3 2
2 1
Case #3: 2
2 3
2 2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
