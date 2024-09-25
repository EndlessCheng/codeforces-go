// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/762/C
// https://codeforces.com/problemset/status/762/problem/C
func Test_cf762C(t *testing.T) {
	testCases := [][2]string{
		{
			`hi
bob`,
			`-`,
		},
		{
			`abca
accepted`,
			`ac`,
		},
		{
			`abacaba
abcdcba`,
			`abcba`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf762C)
}