package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`2 3 4
1 3
2 1
2 2
2 3`,
			`row 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
