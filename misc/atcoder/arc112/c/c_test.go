// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc112/tasks/arc112_c
// 提交：https://atcoder.jp/contests/arc112/submit?taskScreenName=arc112_c
// 对拍：https://atcoder.jp/contests/arc112/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc112_c&orderBy=source_length
// 最短：https://atcoder.jp/contests/arc112/submissions?f.Status=AC&f.Task=arc112_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`10
1 2 3 4 5 6 7 8 9`,
			`10`,
		},
		{
			`5
1 2 3 1`,
			`2`,
		},
		{
			`10
1 1 3 1 3 6 7 6 6`,
			`5`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
