// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/2037/problem/G
// https://codeforces.com/problemset/status/2037/problem/G?friends=on
func Test_cf2037G(t *testing.T) {
	testCases := [][2]string{
		{
			`5
2 6 3 4 6`,
			`5`,
		},
		{
			`5
4 196 2662 2197 121`,
			`2`,
		},
		{
			`7
3 6 8 9 11 12 20`,
			`7`,
		},
		{
			`2
2 3`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2037G)
}