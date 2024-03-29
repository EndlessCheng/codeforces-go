// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/804/B
// https://codeforces.com/problemset/status/804/problem/B
func TestCF804B(t *testing.T) {
	testCases := [][2]string{
		{
			`ab`,
			`1`,
		},
		{
			`aab`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF804B)
}
