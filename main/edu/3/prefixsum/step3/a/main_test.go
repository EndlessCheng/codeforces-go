package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1 1
104
1
1 1 1 1`,
			`104`,
		},
		{
			`2 2
1 2
4 8
7
1 1 1 1
1 1 1 2
1 1 2 1
1 1 2 2
1 2 1 2
2 1 2 1
2 2 2 2`,
			`1
3
5
15
2
4
8`,
		},
		{
			`3 3
1 2 3
-4 -5 -6
7 8 -9
5
1 1 3 3
2 2 2 2
3 1 3 2
1 3 3 3
1 2 3 3`,
			`-3
-5
15
-12
-7`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
