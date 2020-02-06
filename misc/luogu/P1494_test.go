package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p1494(t *testing.T) {
	customInputs := []string{
		`6 4
1 2 3 3 3 2
2 6
1 3
3 5
1 6`,
	}
	customAnswers := []string{
		`2/5
0/1
1/1
4/15`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, p1494)
}
