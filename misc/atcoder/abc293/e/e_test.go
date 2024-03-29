// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc293/submit?taskScreenName=abc293_e
// 对拍：https://atcoder.jp/contests/abc293/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc293_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`3 4 7`,
			`5`,
		},
		{
			`8 10 9`,
			`0`,
		},
		{
			`1000000000 1000000000000 998244353`,
			`919667211`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc293/tasks/abc293_e
