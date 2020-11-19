package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`7
1
3
2
4
5
3
9`,
			`3`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
