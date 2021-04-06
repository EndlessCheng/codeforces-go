package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1510/G
// https://codeforces.com/problemset/status/1510/problem/G
func TestCF1510G(t *testing.T) {
	testCases := [][2]string{
		{
			`3
6 2
1 1 2 2 3
6 6
1 1 2 2 3
6 4
1 2 3 4 5`,
			`1
1 2
8
1 3 6 3 1 2 5 2 4
3
1 2 3 4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF1510G)
}
