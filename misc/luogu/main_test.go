package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`3
1 1
1 2
2 2`,
	}
	customAnswers := []string{
		`1.0000`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
