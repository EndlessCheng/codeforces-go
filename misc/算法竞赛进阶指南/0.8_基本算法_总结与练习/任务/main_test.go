package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`1 2
100 3
100 2
100 1`,
			`1 50004`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
