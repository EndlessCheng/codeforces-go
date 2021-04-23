package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3
1 2 -3 4 5
1 2 3
2 2 -1
1 3 2`,
			`2
-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
