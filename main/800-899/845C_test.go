// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/845/C
// https://codeforces.com/problemset/status/845/problem/C
func Test_cf845C(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 2
2 3
4 5`,
			`YES`,
		},
		{
			`4
1 2
2 3
2 3
1 2`,
			`NO`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf845C)
}
