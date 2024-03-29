// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/601/A
// https://codeforces.com/problemset/status/601/problem/A
func Test_cf601A(t *testing.T) {
	testCases := [][2]string{
		{
			`4 2
1 3
3 4`,
			`2`,
		},
		{
			`4 6
1 2
1 3
1 4
2 3
2 4
3 4`,
			`-1`,
		},
		{
			`5 5
4 2
3 5
4 5
5 1
1 2`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf601A)
}
