// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1627/E
// https://codeforces.com/problemset/status/1627/problem/E?friends=on
func Test_cf1627E(t *testing.T) {
	testCases := [][2]string{
		{
			`4
5 3 3
5 17 8 1 4
1 3 3 3 4
3 1 5 2 5
3 2 5 1 6
6 3 3
5 17 8 1 4 2
1 3 3 3 4
3 1 5 2 5
3 2 5 1 6
5 3 1
5 17 8 1 4
1 3 5 3 100
5 5 5
3 2 3 7 5
3 5 4 2 1
2 2 5 4 5
4 4 5 2 3
1 2 4 2 2
3 3 5 2 4`,
			`16
NO ESCAPE
-90
27`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1627E)
}