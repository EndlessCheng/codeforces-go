// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc185/submit?taskScreenName=abc185_d
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`5 2
1 3`,
			`3`,
		},
		{
			`13 3
13 3 9`,
			`6`,
		},
		{
			`5 5
5 2 1 4 3`,
			`0`,
		},
		{
			`1 0`,
			`1`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc185/tasks/abc185_d
