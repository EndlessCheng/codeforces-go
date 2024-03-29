// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/526/D
// https://codeforces.com/problemset/status/526/problem/D
func TestCF526D(t *testing.T) {
	testCases := [][2]string{
		{
			`7 2
bcabcab`,
			`0000011`,
		},
		{
			`21 2
ababaababaababaababaa`,
			`000110000111111000011`,
		},
		{
			`2 1
ab`,
			`11`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF526D)
}
