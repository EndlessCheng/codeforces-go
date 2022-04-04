package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1660/problem/B
// https://codeforces.com/problemset/status/1660/problem/B
func TestCF1660B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2
2 3
1
2
5
1 6 2 4 3
4
2 2 2 1
3
1 1000000000 999999999
1
1
outputCopy
YES
NO
NO
YES
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1660B)
}
