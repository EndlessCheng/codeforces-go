// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc133/submit?taskScreenName=arc133_c
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`2 4 3
0 2
1 2 2 0`,
			`11`,
		},
		{
			`3 3 4
0 1 2
1 2 3`,
			`-1`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc133/tasks/arc133_c
