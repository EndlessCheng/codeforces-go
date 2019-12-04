package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`5 5 4
3 1
2 4
5 1
1 4
2 4
3 2
3 5
1 2
4 5`,
	}
	customAnswers := []string{
		`4
4
1
4
4`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
