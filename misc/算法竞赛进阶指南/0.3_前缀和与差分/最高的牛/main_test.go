package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`9 3 5 5
1 3
5 3
4 3
3 7
9 8`,
			`5
4
5
3
4
4
5
5
5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
