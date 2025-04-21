package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://loj.ac/p/160
func Test_p160(t *testing.T) {
	testCases := [][2]string{
		{
			`7 0
3 4 5 3 0 7 0
1 2 2 1 1 2 2
1 1 2 3 1 1 2`,
			`0`,
		},
		{
			`7 4
3 4 5 3 0 7 0
1 2 2 1 1 2 2
1 1 2 3 1 1 2`,
			`6`,
		},
		{
			`10 6
0 1 1 1 2 3 4 5 6 7
1 1 2 2 3 1 2 2 3 1
0 1 1 1 1 1 1 1 1 1`,
			`3`,
		},
		{
			`13 3
0 1 1 1 3 3 0 0 8 9 10 8 10
1 1 1 1 1 1 1 1 1 1 1 1 1
2 16 32 512 2048 4096 4 1 8 64 1024 128 256`,
			`4130`,
		},
		{
			`13 25
0 1 1 1 3 3 0 0 8 9 10 8 10
5 6 4 3 7 5 4 7 3 6 5 6 6
0 2 1 2 3 3 3 0 1 2 4 3 2`,
			`10`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p160)
}
