// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`1 1
1`,
			`2`,
		},
		{
			`1 1
10`,
			`20`,
		},
		{
			`2 4
10 10`,
			`80`,
		},
		{
			`3 1
10 10 10`,
			`20`,
		},
		{
			`5 3
10 14 19 34 33`,
			`202`,
		},
		{
			`9 14
1 3 5 110 24 21 34 5 3`,
			`1837`,
		},
		{
			`9 73
67597 52981 5828 66249 75177 64141 40773 79105 16076`,
			`8128170`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}

// https://atcoder.jp/contests/abc149/tasks/abc149_e
// https://atcoder.jp/contests/abc149/submit?taskScreenName=abc149_e
