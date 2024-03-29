// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc125/tasks/abc125_d
// 提交：https://atcoder.jp/contests/abc125/submit?taskScreenName=abc125_d
// 对拍：https://atcoder.jp/contests/abc125/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc125_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`3
-10 5 -4`,
			`19`,
		},
		{
			`5
10 -4 -8 -11 3`,
			`30`,
		},
		{
			`11
-1000000000 1000000000 -1000000000 1000000000 -1000000000 0 1000000000 -1000000000 1000000000 -1000000000 1000000000`,
			`10000000000`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
