package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3 20
1 2 3`,
			`2
2 9 9`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
