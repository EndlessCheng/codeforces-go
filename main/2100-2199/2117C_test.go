// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2117/C
// https://codeforces.com/problemset/status/2117/problem/C?friends=on
func Test_cf2117C(t *testing.T) {
	testCases := [][2]string{
		{
			`8
6
1 2 2 3 1 5
8
1 2 1 3 2 1 3 2
5
5 4 3 2 1
10
5 8 7 5 8 5 7 8 10 9
3
1 2 2
9
3 3 1 4 3 2 4 1 2
6
4 5 4 5 6 4
8
1 2 1 2 1 2 1 2`,
			`2
3
1
3
1
3
3
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2117C)
}
