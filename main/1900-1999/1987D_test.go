// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1987/problem/D
// https://codeforces.com/problemset/status/1987/problem/D
func Test_cf1987D(t *testing.T) {
	testCases := [][2]string{
		{
			`9
4
1 4 2 3
3
1 1 1
5
1 4 2 3 4
4
3 4 1 4
1
1
8
4 3 2 5 6 8 3 4
7
6 1 1 3 5 3 1
11
6 11 6 8 7 5 3 11 2 3 5
17
2 6 5 3 9 1 6 2 5 6 3 2 3 9 6 1 6`,
			`2
1
3
2
1
3
2
4
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1987D)
}
