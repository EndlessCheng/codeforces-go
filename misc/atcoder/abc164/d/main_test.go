package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	testCases := [][2]string{
		{
			`1817181712114`,
			`3`,
		},
		{
			`14282668646`,
			`2`,
		},
		{
			`2119`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
