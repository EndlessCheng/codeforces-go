// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/B
// https://codeforces.com/problemset/status/1037/problem/B?friends=on
func Test_cf1037B(t *testing.T) {
	testCases := [][2]string{
		{
			`3 8
6 5 8`,
			`2`,
		},
		{
			`7 20
21 15 12 11 20 19 12`,
			`6`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1037B)
}
