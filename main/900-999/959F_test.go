// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/959/F
// https://codeforces.com/problemset/status/959/problem/F
func Test_cf959F(t *testing.T) {
	testCases := [][2]string{
		{
			`5 5
0 1 2 3 4
4 3
2 0
3 7
5 7
5 8`,
			`4
2
0
4
0`,
		},
		{
			`3 2
1 1 1
3 1
2 0`,
			`4
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf959F)
}
