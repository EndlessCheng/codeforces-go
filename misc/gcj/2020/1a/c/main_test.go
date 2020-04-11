package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	customInputs := []string{
		`4
1 1
15
3 3
1 1 1
1 2 1
1 1 1
1 3
3 1 2
1 3
1 2 3`,
	}
	customAnswers := []string{
		`Case #1: 15
Case #2: 16
Case #3: 14
Case #4: 14`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
}
