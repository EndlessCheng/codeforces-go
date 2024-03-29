// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1550/E
// https://codeforces.com/problemset/status/1550/problem/E
func Test_cf1550E(t *testing.T) {
	testCases := [][2]string{
		{
			`10 2
a??ab????b`,
			`4`,
		},
		{
			`9 4
?????????`,
			`2`,
		},
		{
			`2 3
??`,
			`0`,
		},
		{
			`15 3
??b?babbc??b?aa`,
			`3`,
		},
		{
			`4 4
cabd`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1550E)
}
