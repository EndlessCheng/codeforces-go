// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc076/tasks/arc076_b
// 提交：https://atcoder.jp/contests/arc076/submit?taskScreenName=arc076_b
// 对拍：https://atcoder.jp/contests/arc076/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc076_b&orderBy=source_length
// 最短：https://atcoder.jp/contests/arc076/submissions?f.Status=AC&f.Task=arc076_b&orderBy=source_length
func Test_b(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 5
3 9
7 8`,
			`3`,
		},
		{
			`6
8 3
4 9
12 19
18 1
13 5
7 6`,
			`8`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
