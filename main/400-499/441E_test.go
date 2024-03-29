// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/441/E
// https://codeforces.com/problemset/status/441/problem/E
func Test_cf441E(t *testing.T) {
	testCases := [][2]string{
		{
			`1 1 50`,
			`1.0000000000000`,
		},
		{
			`5 3 0`,
			`3.0000000000000`,
		},
		{
			`5 3 25`,
			`1.9218750000000`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf441E)
}
