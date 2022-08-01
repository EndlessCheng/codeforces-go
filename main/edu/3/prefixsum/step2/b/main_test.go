package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`2
12 6
3
1 1
2 2
1 2`,
			`12
6
10`,
		},
		{
			`5
1 2 3 0 4
5
1 5
1 2
4 4
3 4
2 5`,
			`4
3
0
3
5`,
		},
		{
			`7
1 1 1 1 1 1 1
9
1 7
2 6
3 4
2 5
5 7
6 6
3 3
2 7
1 1`,
			`1
1
0
0
1
1
1
0
1`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
