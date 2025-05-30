// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc151/tasks/abc151_e
// 提交：https://atcoder.jp/contests/abc151/submit?taskScreenName=abc151_e
// 对拍：https://atcoder.jp/contests/abc151/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc151_e&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc151/submissions?f.Status=AC&f.Task=abc151_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`4 2
1 1 3 4`,
			`11`,
		},
		{
			`6 3
10 10 10 -10 -10 -10`,
			`360`,
		},
		{
			`3 1
1 1 1`,
			`0`,
		},
		{
			`10 6
1000000000 1000000000 1000000000 1000000000 1000000000 0 0 0 0 0`,
			`999998537`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
