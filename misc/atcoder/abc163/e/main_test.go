package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1 3 4 2`,
			`20`,
		},
		{
			`6
5 5 6 1 1 1`,
			`58`,
		},
		{
			`6
8 6 9 1 2 1`,
			`85`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
