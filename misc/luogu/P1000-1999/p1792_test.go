package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p1792(t *testing.T) {
	testCases := [][2]string{
		{
			`7 3
1 2 3 4 5 6 7`,
			`15`,
		},
		{
			`7 4
1 2 3 4 5 6 7`,
			`Error!`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1792)
}
