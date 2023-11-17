package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test1484(t *testing.T) {
	cases := [][2]string{
		{
			`6 3 
100 1 -1 100 1 -1`,
			`200`,
		},
	}
	tarCase := 0
	testutil.AssertEqualStringCase(t, cases, tarCase, run1484)
}
