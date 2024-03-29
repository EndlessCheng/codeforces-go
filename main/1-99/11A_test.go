// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/11/A
// https://codeforces.com/problemset/status/11/problem/A
func Test_cf11A(t *testing.T) {
	testCases := [][2]string{
		{
			`4 2
1 3 3 2`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf11A)
}
