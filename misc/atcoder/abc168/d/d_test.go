// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`4 4
1 2
2 3
3 4
4 2`,
			`Yes
1
2
2`,
		},
		{
			`6 9
3 4
6 1
2 4
5 3
4 6
1 5
6 2
4 5
5 6`,
			`Yes
6
5
5
1
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc168/tasks/abc168_d
