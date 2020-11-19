package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1
30
5
1 1 1 1 1
3
3 2 1
4
5 3 3 1
5
6 5 4 3 2
2
15 15
0`,
			`1
1
16
4158
141892608
9694845`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
