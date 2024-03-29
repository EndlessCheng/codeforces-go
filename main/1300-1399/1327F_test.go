// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1327/F
// https://codeforces.com/problemset/status/1327/problem/F
func TestCF1327F(t *testing.T) {
	testCases := [][2]string{
		{
			`4 3 2
1 3 3
3 4 6`,
			`3`,
		},
		{
			`5 2 3
1 3 2
2 5 0
3 3 3`,
			`33`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1327F)
}
