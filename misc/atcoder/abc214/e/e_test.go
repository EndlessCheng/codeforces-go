// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc214/submit?taskScreenName=abc214_e
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`2
3
1 2
2 3
3 3
5
1 2
2 3
3 3
1 3
999999999 1000000000`,
			`Yes
No`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc214/tasks/abc214_e
