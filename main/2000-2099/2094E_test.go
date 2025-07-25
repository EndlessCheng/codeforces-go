// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2094/E
// https://codeforces.com/problemset/status/2094/problem/E?friends=on
func Test_cf2094E(t *testing.T) {
	testCases := [][2]string{
		{
			`5
3
18 18 18
5
1 2 4 8 16
5
8 13 4 5 15
6
625 676 729 784 841 900
1
1`,
			`0
79
37
1555
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2094E)
}
