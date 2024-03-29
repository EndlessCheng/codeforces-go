// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc194/tasks/abc194_e
// 提交：https://atcoder.jp/contests/abc194/submit?taskScreenName=abc194_e
// 对拍：https://atcoder.jp/contests/abc194/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc194_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3 2
0 0 1`,
			`1`,
		},
		{
			`3 2
1 1 1`,
			`0`,
		},
		{
			`3 2
0 1 0`,
			`2`,
		},
		{
			`7 3
0 0 1 2 0 1 0`,
			`2`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
