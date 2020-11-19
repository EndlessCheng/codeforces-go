package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3
2
3
187`,
			`1666
2666
66666`,
		},
		{
			`1
50000000`,
			`6668056399`,
		},
		{
			`2
0
1`,
			``,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
