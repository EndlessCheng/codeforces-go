package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p4447(t *testing.T) {
	samples := [][2]string{
		{
			`7
4 5 2 3 -4 -3 -5`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, samples, 0, p4447)
}
