package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`1 2
1 2
1 2 1`,
	}
	customAnswers := []string{
		`1 4 5 2`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
