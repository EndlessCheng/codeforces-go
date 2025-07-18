// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2106/D
// https://codeforces.com/problemset/status/2106/problem/D?friends=on
func Test_cf2106D(t *testing.T) {
	testCases := [][2]string{
		{
			`7
9 5
3 5 2 3 3 5 8 1 2
4 6 2 4 6
6 3
1 2 6 8 2 1
5 4 3
5 3
4 3 5 4 3
7 4 5
6 3
8 4 2 1 2 5
6 1 4
5 5
1 2 3 4 5
5 4 3 2 1
6 3
1 2 3 4 5 6
9 8 7
5 5
7 7 6 7 7
7 7 7 7 7`,
			`6
3
7
0
-1
-1
7`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2106D)
}
