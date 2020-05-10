package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`2 1 1 3`,
			`2`,
		},
		{
			`1 2 3 4`,
			`0`,
		},
		{
			`2000000000 0 0 2000000000`,
			`2000000000`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc167/tasks/abc167_b
