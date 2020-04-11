package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	customInputs := []string{
		``,
	}
	customAnswers := []string{
		``,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
	//testutil.AssertEqualRunResults(t, customInputs, 0, runAC, run)
	_ = customAnswers
}
