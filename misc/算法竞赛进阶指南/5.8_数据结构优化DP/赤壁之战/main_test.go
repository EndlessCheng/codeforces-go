package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`2
3 2
1 2 3
3 2
3 2 1`,
			`Case #1: 3
Case #2: 0`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
