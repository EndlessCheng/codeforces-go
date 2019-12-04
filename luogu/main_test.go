package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`3 3
1 2
2 1
2 3`,
	}
	customAnswers := []string{
		`1`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
