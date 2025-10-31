package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p2895(t *testing.T) {
	cases := [][2]string{
		{
			`4
0 0 2
2 1 2
1 1 2
0 3 5`,
			`5`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, p2895)
}
