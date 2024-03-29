// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc201/submit?taskScreenName=abc201_e
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3
1 2 1
1 3 3`,
			`6`,
		},
		{
			`5
3 5 2
2 3 2
1 5 1
4 5 13`,
			`62`,
		},
		{
			`10
5 7 459221860242673109
6 8 248001948488076933
3 5 371922579800289138
2 5 773108338386747788
6 10 181747352791505823
1 3 803225386673329326
7 8 139939802736535485
9 10 657980865814127926
2 4 146378247587539124`,
			`241240228`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc201/tasks/abc201_e
