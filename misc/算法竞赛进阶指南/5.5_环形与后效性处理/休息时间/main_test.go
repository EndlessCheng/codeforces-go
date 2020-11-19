package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`5 3
2
0
3
1
4`,
			`6`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
