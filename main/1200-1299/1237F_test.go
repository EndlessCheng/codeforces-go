// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1237/F
// https://codeforces.com/problemset/status/1237/problem/F
func Test_cf1237F(t *testing.T) {
	testCases := [][2]string{
		{
			`5 7 2
3 1 3 2
4 4 4 5`,
			`8`,
		},
		{
			`5 4 2
1 2 2 2
4 3 4 4`,
			`1`,
		},
		{
			`23 42 0`,
			`102848351`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1237F)
}
