package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p3373(t *testing.T) {
	samples := [][2]string{
		{
			`5 5 571373
1 5 4 2 3
2 1 4 1
3 2 5
1 2 4 2
2 3 5 5
3 1 4`,
			`17
40`,
		},
	}
	testutil.AssertEqualStringCase(t, samples, 0, p3373)
}
