package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p2586(t *testing.T) {
	customInputs := []string{
		`3 5
1 1 2
2 2
5`,
	}
	customAnswers := []string{
		`The game is going on
5
5 1 3 1 4
4 1 3 0 4
3 1 3 0 3
2 1 3 0 2
1 1 4 0 1`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, p2586)
}
