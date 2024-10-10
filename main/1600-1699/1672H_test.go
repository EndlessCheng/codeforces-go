// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1672/H
// https://codeforces.com/problemset/status/1672/problem/H
func Test_cf1672H(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3
11011
2 4
1 5
3 5`,
			`1
3
2`,
		},
		{
			`10 3
1001110110
1 10
2 5
5 10`,
			`4
2
3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1672H)
}
