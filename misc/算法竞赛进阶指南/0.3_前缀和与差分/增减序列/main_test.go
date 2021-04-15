package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1
1
2
2`,
			`1
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
