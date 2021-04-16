package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`2
5 1 49
8 2 1 7 9
5 1 64
8 2 1 7 9`,
			`2
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, -1, run)
}
