package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 1
2 3
7 4
4 6`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
