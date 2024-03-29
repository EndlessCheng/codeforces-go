// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc191/submit?taskScreenName=abc191_d
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`0.2 0.8 1.1`,
			`3`,
		},
		{
			`100 100 1`,
			`5`,
		},
		{
			`42782.4720 31949.0192 99999.99`,
			`31415920098`,
		},
		// TODO 测试参数的下界和上界
		{
			`0.0001 100000 100000`,
			`31415925440`,
		},
		{
			`0.0001 93819 93819`,
			`27652312476`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc191/tasks/abc191_d
