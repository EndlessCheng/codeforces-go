package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	customInputs := []string{
		`5
3
501
512
1000
1000000000`,
	}
	customAnswers := []string{
		`Case #1:
1 1
Case #2:
1 1
2 1
2 2
3 3
Case #3:
1 1
2 2
3 2
4 3
5 3
5 2
4 1
3 1`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
}
