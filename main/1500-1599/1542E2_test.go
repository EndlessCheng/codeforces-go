// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1542/E2
// https://codeforces.com/problemset/status/1542/problem/E2?friends=on
func Test_cf1542E2(t *testing.T) {
	testCases := [][2]string{
		{
			`4 403458273`,
			`17`,
		},
		{
			`50 1000000000`,
			`14460084`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1542E2)
}
