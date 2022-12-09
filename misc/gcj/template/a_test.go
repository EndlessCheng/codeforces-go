package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var customTestCases = [][2]string{
	{
		``,
		``,
	},
}

func Test_a(t *testing.T) {
	tar := 0 // -1
	testutil.AssertEqualStringCase(t, customTestCases, tar, run)
}
