// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1526/E
// https://codeforces.com/problemset/status/1526/problem/E?friends=on
func Test_cf1526E(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2
0 2 1`,
			`1`,
		},
		{
			`5 1
0 1 2 3 4`,
			`0`,
		},
		{
			`6 200000
0 1 2 3 4 5`,
			`822243495`,
		},
		{
			`7 6
3 2 4 1 0 5 6`,
			`36`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1526E)
}
