package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p2850(t *testing.T) {
	cases := [][2]string{
		{
			`2
3 3 1
1 2 2
1 3 4
2 3 1
3 1 3
3 2 1
1 2 3
2 3 4
3 1 8`,
			`NO
YES`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, p2850)
}
