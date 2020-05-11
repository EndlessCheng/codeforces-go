package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`6
(((
()))(
(((()
)))
(()))))((((((((((
(())))))))))(((`,
			``,
		},
		{
			`2
)
(()`,
			`Yes`,
		},
		{
			`2
)(
()`,
			`No`,
		},
		{
			`4
((()))
((((((
))))))
()()()`,
			`Yes`,
		},
		{
			`3
(((
)
)`,
			`No`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc167/tasks/abc167_f
