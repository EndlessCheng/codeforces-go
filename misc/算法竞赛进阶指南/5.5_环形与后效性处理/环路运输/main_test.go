package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`5
1 8 6 2 5`,
			`15`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
