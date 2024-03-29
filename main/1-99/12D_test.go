// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/12/D
// https://codeforces.com/problemset/status/12/problem/D
func Test_cf12D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 4 2
4 3 2
2 5 3`,
			`1`,
		},
		{
			`5
2 8 10 0 7
7 7 3 0 10
2 8 3 2 2`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf12D)
}
