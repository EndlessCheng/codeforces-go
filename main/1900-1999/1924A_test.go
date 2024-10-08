// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1924/A
// https://codeforces.com/problemset/status/1924/problem/A
func Test_cf1924A(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2 2 4
abba
2 2 3
abb
3 3 10
aabbccabab`,
			`YES
NO
aa
NO
ccc`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1924A)
}
