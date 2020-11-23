package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	cases := [][2]string{
		{
			`3 4 2 
2 2 
2 3`,
			`2`,
		},
		{
			`100 100 3
15 16
16 15
99 88`,
			`545732279`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, run)
}
