package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1613/problem/C
// https://codeforces.com/problemset/status/1613/problem/C
func TestCF1613C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 5
1 5
3 10
2 4 10
5 3
1 2 4 5 7
4 1000
3 25 64 1337
outputCopy
3
4
1
470`
	testutil.AssertEqualCase(t, rawText, 0, CF1613C)
}
