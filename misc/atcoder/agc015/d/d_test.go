// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/agc015/tasks/agc015_d
// 提交：https://atcoder.jp/contests/agc015/submit?taskScreenName=agc015_d
// 对拍：https://atcoder.jp/contests/agc015/submissions?f.LanguageName=Go&f.Status=AC&f.Task=agc015_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`7
9`,
			`4`,
		},
		{
			`65
98`,
			`63`,
		},
		{
			`271828182845904523
314159265358979323`,
			`68833183630578410`,
		},
		{
			`69946535201660593
69964127387705005`,
			`35180883026255`,
		},
		{
			`126
252`,
			`130`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
