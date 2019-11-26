package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`5 5
25957 6405 15770 26287 26465 
2 2 1
3 4 1
4 5 1
1 2 2
4 4 1`,
	}
	customAnswers := []string{
		`6405
15770
26287
25957
26287`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
