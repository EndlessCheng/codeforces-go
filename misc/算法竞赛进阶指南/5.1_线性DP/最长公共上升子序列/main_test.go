package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`4
2 2 1 3
2 1 2 3`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
