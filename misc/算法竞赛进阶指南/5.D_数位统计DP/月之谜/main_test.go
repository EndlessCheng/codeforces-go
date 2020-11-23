package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`1 100`,
			`33`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
