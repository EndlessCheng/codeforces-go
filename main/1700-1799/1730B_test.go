// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1730/B
// https://codeforces.com/problemset/status/1730/problem/B
func Test_cf1730B(t *testing.T) {
	testCases := [][2]string{
		{
			`7
1
0
3
2
3 1
0 0
2
1 4
0 0
3
1 2 3
0 0 0
3
1 2 3
4 1 2
3
3 3 3
5 3 3
6
5 4 7 2 10 4
3 2 5 1 4 6`,
			`0
2
2.5
2
1
3
6`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1730B)
}
