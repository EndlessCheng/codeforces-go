// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc171/tasks/abc171_e
// 提交：https://atcoder.jp/contests/abc171/submit?taskScreenName=abc171_e
// 对拍：https://atcoder.jp/contests/abc171/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc171_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`4
20 11 9 24`,
			`26 5 7 22`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
