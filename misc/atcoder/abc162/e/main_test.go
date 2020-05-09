package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2`,
			`9`,
		},
		{
			`3 200`,
			`10813692`,
		},
		{
			`100000 100000`,
			`742202979`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc162/tasks/abc162_e
