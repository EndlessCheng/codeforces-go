package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3 3 10
60 2 2 4
70 8 7 9
50 2 3 9`,
			`120`,
		},
		{
			`3 3 10
100 3 1 4
100 1 5 9
100 2 6 5`,
			`-1`,
		},
		{
			`8 5 22
100 3 7 5 3 1
164 4 5 2 7 8
334 7 2 7 2 9
234 4 7 2 8 2
541 5 4 3 3 6
235 4 8 6 9 7
394 3 6 1 6 2
872 8 4 3 7 2`,
			`1067`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc167/tasks/abc167_c
