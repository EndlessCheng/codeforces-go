// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1609/E
// https://codeforces.com/problemset/status/1609/problem/E?friends=on
func Test_cf1609E(t *testing.T) {
	testCases := [][2]string{
		{
			`9 12
aaabccccc
4 a
4 b
2 b
5 a
1 b
6 b
5 c
2 a
1 a
5 a
6 b
7 b`,
			`0
1
2
2
1
2
1
2
2
2
2
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1609E)
}
