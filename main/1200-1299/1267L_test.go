// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1267/L
// https://codeforces.com/problemset/status/1267/problem/L
func Test_cf1267L(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2 2
abcdef`,
			`af
bc
ed`,
		},
		{
			`2 3 1
abcabc`,
			`aab
bcc`,
		},
		{
			`2 3 2
bbcadc`,
			`acd
bbc`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1267L)
}
