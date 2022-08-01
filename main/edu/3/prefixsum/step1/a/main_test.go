package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1
178`,
			`0 178 `,
		},
		{
			`7
0 1 2 3 4 5 6`,
			`0 0 1 3 6 10 15 21 `,
		},
		{
			`5
-1 -2 3 0 4`,
			`0 -1 -3 0 0 4 `,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
