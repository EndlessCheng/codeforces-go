// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1822/problem/F
// https://codeforces.com/problemset/status/1822/problem/F
func Test_cf1822F(t *testing.T) {
	testCases := [][2]string{
		{
			`4
3 2 3
2 1
3 1
5 4 1
2 1
4 2
5 4
3 4
6 5 3
4 1
6 1
2 6
5 1
3 2
10 6 4
1 3
1 9
9 7
7 6
6 4
9 2
2 8
8 5
5 10`,
			`2
12
17
32`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1822F)
}
