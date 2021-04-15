package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`3 10
AND 5
OR 6
XOR 7`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
