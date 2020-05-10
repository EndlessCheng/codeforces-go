package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	testCases := [][2]string{
		{
			`4 1`,
			`1 4`,
		},
		{
			`7 3`,
			`1 7
2 6
3 5`,
		},
		{
			`6 2`,
			`1 6
2 4`,
		},
		{
			`8 3`,
			`1 8
2 7
3 5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc165/tasks/abc165_e
