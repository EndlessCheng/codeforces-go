// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/17/E
// https://codeforces.com/problemset/status/17/problem/E?friends=on
func Test_cf17E(t *testing.T) {
	testCases := [][2]string{
		{
			`4
babb`,
			`6`,
		},
		{
			`2
aa`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf17E)
}
