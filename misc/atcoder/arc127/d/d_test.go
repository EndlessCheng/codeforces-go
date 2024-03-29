// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc127/tasks/arc127_d
// 提交：https://atcoder.jp/contests/arc127/submit?taskScreenName=arc127_d
// 对拍：https://atcoder.jp/contests/arc127/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc127_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`3
1 2 3
4 5 6`,
			`4`,
		},
		{
			`4
1 2 3 4
1 2 3 4`,
			`24`,
		},
		{
			`10
195247 210567 149398 9678 23694 46151 187762 17915 176476 249828
68649 128425 249346 62366 194119 117620 26327 161384 207 57656`,
			`4019496`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
