package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3 10
1 2 4 2 1 1
2 5
1 4 2 1
0 0`,
			`8
4`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
