// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1837/problem/F
// https://codeforces.com/problemset/status/1837/problem/F
func Test_cf1837F(t *testing.T) {
	testCases := [][2]string{
		{
			`6
5 4
1 10 1 1 1
5 3
1 20 5 15 3
5 3
1 20 3 15 5
10 6
10 8 20 14 3 8 6 4 16 11
10 5
9 9 2 13 15 19 4 9 13 12
1 1
1`,
			`2
6
5
21
18
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1837F)
}
