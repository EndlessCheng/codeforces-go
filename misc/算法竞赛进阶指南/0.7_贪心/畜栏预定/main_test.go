package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 10
2 4
3 6
5 8
4 7`,
			`4
1
2
3
2
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
