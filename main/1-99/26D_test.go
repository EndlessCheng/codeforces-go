// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/26/D
// https://codeforces.com/problemset/status/26/problem/D?friends=on
func Test_cf26D(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3 1`,
			`0.857143`,
		},
		{
			`0 5 5`,
			`1`,
		},
		{
			`0 1 0`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf26D)
}
