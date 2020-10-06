package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p2868(t *testing.T) {
	cases := [][2]string{
		{
			`5 7
30
10
10
5
10
1 2 3
2 3 2
3 4 5
3 5 2
4 5 5
5 1 3
5 2 2`,
			`6.00`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, p2868)
}
