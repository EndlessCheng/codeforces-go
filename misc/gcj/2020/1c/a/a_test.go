package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`5
4 4 SSSS
3 0 SNSS
2 10 NSNNSN
0 1 S
2 7 SSSSSSSS`,
			`Case #1: 4
Case #2: IMPOSSIBLE
Case #3: IMPOSSIBLE
Case #4: 1
Case #5: 5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
