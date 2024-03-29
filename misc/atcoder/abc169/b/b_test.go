// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`2
1000000000 1000000000`,
			`1000000000000000000`,
		},
		{
			`3
101 9901 999999000001`,
			`-1`,
		},
		{
			`31
4 1 5 9 2 6 5 3 5 8 9 7 9 3 2 3 8 4 6 2 6 4 3 3 8 3 2 7 9 5 0`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc169/tasks/abc169_b
// https://atcoder.jp/contests/abc169/submit?taskScreenName=abc169_b
