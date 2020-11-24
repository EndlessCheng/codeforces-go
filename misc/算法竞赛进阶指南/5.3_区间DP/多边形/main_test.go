package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`4
t -7 t 4 x 2 x 5`,
			`33
1 2`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
