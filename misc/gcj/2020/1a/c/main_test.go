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
1 2 3`, `5
4 4
1 2 1 2
2 1 2 1
1 2 1 2
2 1 2 1
4 4
1 2 1 2
2 2 2 1
1 2 1 2
2 1 2 1
4 4
1 2 1 2
2 2 2 1
1 2 2 2
2 1 2 1
4 4
1 2 1 3
2 3 2 1
1 2 1 2
2 1 2 1
4 4
1 2 1 1
2 1 1 1
1 2 1 2
2 1 2 1`,
	}
	customAnswers := []string{
		`Case #1: 15
Case #2: 16
Case #3: 14
Case #4: 14`, `Case #1: 40
Case #2: 43
Case #3: 46
Case #4: 57
Case #5: 48`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
	//testutil.AssertEqualRunResults(t, customInputs, 1, runAC, run)
	_ = customAnswers
}
