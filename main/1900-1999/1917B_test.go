// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1917/problem/B
// https://codeforces.com/problemset/status/1917/problem/B
func Test_cf1917B(t *testing.T) {
	testCases := [][2]string{
		{
			`5
5
aaaaa
1
z
5
ababa
14
bcdaaaabcdaaaa
20
abcdefghijklmnopqrst`,
			`5
1
9
50
210`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1917B)
}
