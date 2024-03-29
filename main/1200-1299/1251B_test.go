// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1251/problem/B
// https://codeforces.com/problemset/status/1251/problem/B
func Test_cf1251B(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1
0
3
1110
100110
010101
2
11111
000001
2
001
11100111`,
			`1
2
2
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1251B)
}
