package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF102426J(t *testing.T) {
	testCases := [][2]string{
		{
			`5`,
			`13`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF102426J)
}
