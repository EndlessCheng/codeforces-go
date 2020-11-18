package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1 2
1 3
1 4
2 2
2 3
2 4
2 11
4 11
0 0`,
			`1
0
1
2
3
5
144
51205`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
