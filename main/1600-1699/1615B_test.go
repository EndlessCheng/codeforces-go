package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1615/B
// https://codeforces.com/problemset/status/1615/problem/B
func TestCF1615B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
2 8
4 5
1 5
100000 200000
outputCopy
1
3
0
2
31072`
	testutil.AssertEqualCase(t, rawText, 0, CF1615B)
}
