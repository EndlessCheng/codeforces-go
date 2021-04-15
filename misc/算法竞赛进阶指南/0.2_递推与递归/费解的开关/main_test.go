package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`3
00111
01011
10001
11010
11100

11101
11101
11110
11111
11111

01111
11111
11111
11111
11111`,
			`3
2
-1`,
		},
		{
			`1
11101
11101
11110
11111
11111`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, -1, run)
}
