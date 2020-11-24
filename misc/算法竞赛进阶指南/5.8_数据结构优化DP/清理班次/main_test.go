package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3 0 4
0 2 3
3 4 2
0 0 1`,
			`5`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
