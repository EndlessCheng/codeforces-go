// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc091/tasks/arc092_b
// 提交：https://atcoder.jp/contests/arc092/submit?taskScreenName=arc092_b
// 对拍：https://atcoder.jp/contests/arc092/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc092_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`2
1 2
3 4`,
			`2`,
		},
		{
			`6
4 6 0 0 3 3
0 5 6 5 0 3`,
			`8`,
		},
		{
			`5
1 2 3 4 5
1 2 3 4 5`,
			`2`,
		},
		{
			`1
0
0`,
			`0`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
