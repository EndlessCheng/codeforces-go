package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	customInputs := []string{
		``,
	}
	customAnswers := []string{
		``,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
}
