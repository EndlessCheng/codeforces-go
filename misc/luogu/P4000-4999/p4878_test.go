package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p4878(t *testing.T) {
	cases := [][2]string{
		{
			`4 2 1
1 3 10
2 4 20
2 3 3`,
			`27`,
		},

	}
	testutil.AssertEqualStringCase(t, cases, 0, p4878)
}
