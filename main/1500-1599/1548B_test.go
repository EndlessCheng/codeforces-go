package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1548/problem/B
// https://codeforces.com/problemset/status/1548/problem/B
func TestCF1548B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 5 2 4 6
4
8 2 5 10
2
1000 2000
8
465 55 3 54 234 12 45 78
outputCopy
3
3
2
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1548B)
}
