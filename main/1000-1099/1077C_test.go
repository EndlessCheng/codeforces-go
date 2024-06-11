// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1077/C
// https://codeforces.com/problemset/status/1077/problem/C
func Test_cf1077C(t *testing.T) {
	testCases := [][2]string{
		{
			`5
2 5 1 2 2`,
			`3
4 1 5`,
		},
		{
			`4
8 3 5 2`,
			`2
1 4`,
		},
		{
			`5
2 1 2 4 3`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1077C)
}