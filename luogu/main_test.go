package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		`5 10
59 46 14 87 41
0 2 1
0 1 1 14
0 1 1 57
0 1 1 88
4 2 4
0 2 5
0 2 4
4 2 1
2 2 2
1 1 5 91`,
		`4 5
-869513306 -55515509 -18235282 -850474412 
0 1 1 -371445967
1 1 4 -455963593
0 1 1 -437838418
3 1 3 -212862770
0 2 1`,
	}
	customAnswers := []string{
		`59
87
41
87
88
46`,
		`-869513306`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
}
