// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc147/tasks/abc147_d
// 提交：https://atcoder.jp/contests/abc147/submit?taskScreenName=abc147_d
// 对拍：https://atcoder.jp/contests/abc147/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc147_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`3
1 2 3`,
			`6`,
		},
		{
			`10
3 1 4 1 5 9 2 6 5 3`,
			`237`,
		},
		{
			`10
3 14 159 2653 58979 323846 2643383 27950288 419716939 9375105820`,
			`103715602`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
