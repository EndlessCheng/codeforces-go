package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3 3
0 3 9
2 8 5
5 7 0`,
			`34`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
