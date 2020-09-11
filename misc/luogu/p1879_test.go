package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p1879(t *testing.T) {
	cases := [][2]string{
		{
			`2 3
1 1 1
0 1 0`,
			`9`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, p1879)
}
