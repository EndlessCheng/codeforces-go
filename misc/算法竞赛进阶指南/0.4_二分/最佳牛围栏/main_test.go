package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`10 6
6
4
2
10
3
8
5
9
4
1`,
			`6500`,
		},
		{
			`10 10
6
4
2
10
3
8
5
9
4
1`,
			`5200`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
