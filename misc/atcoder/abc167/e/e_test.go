package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3 2 1`,
			`6`,
		},
		{
			`100 100 0`,
			`73074801`,
		},
		{
			`60522 114575 7559`,
			`479519525`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc167/tasks/abc167_e
