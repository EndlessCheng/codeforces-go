// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc176/tasks/abc176_e
// 提交：https://atcoder.jp/contests/abc176/submit?taskScreenName=abc176_e
// 对拍：https://atcoder.jp/contests/abc176/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc176_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`2 3 3
2 2
1 1
1 3`,
			`3`,
		},
		{
			`3 3 4
3 3
3 1
1 1
1 2`,
			`3`,
		},
		{
			`5 5 10
2 5
4 3
2 3
5 5
2 2
5 4
5 3
5 1
3 5
1 4`,
			`6`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
