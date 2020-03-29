package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	customInputs := []string{
		`100
1`, `25
2`, `314159
2`, `9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
3`,
	}
	customAnswers := []string{
		`19`, `14`, `937`, `117879300`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
}
