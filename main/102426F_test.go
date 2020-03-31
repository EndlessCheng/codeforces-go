package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF102426F(t *testing.T) {
	customInputs := []string{
		`4 3 6
2 2 1
2 2 1
2 2 1
1 1 1`, `3 3 15
1 2 3
4 5 6
7 8 9`,
	}
	customAnswers := []string{
		`1`, `4`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, CF102426F)
}
