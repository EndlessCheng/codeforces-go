package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF102426J(t *testing.T) {
	customInputs := []string{
		`5`,
	}
	customAnswers := []string{
		`13`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, CF102426J)
}
