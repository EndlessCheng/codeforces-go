package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p1948(t *testing.T) {
	customTestCases := [][2]string{
		{
			`5 7 1
1 2 5
3 1 4
2 4 8
3 2 3
5 2 9
3 4 7
4 5 6`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, p1948)
}
