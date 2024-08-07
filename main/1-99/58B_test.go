// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/58/B
// https://codeforces.com/problemset/status/58/problem/B
func Test_cf58B(t *testing.T) {
	testCases := [][2]string{
		{
			`10`,
			`10 5 1`,
		},
		{
			`4`,
			`4 2 1`,
		},
		{
			`3`,
			`3 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf58B)
}
