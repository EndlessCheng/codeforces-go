// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc326/tasks/abc326_d
// 提交：https://atcoder.jp/contests/abc326/submit?taskScreenName=abc326_d
// 对拍：https://atcoder.jp/contests/abc326/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc326_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc326/submissions?f.Status=AC&f.Task=abc326_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`5
ABCBC
ACAAB`,
			`Yes
AC..B
.BA.C
C.BA.
BA.C.
..CBA`,
		},
		{
			`3
AAA
BBB`,
			`No`,
		},
		{
			`4
CCAB
CBCA`,
			`No`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
