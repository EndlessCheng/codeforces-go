package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			``,
			``,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
