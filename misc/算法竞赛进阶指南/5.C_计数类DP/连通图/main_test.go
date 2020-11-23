package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1
2
3
4
0`,
			`1
1
4
38`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
