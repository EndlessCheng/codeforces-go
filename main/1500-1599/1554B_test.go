package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1554/problem/B
// https://codeforces.com/problemset/status/1554/problem/B
func TestCF1554B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 3
1 1 3
2 2
1 2
4 3
0 1 2 3
6 6
3 2 0 0 5 6
outputCopy
-1
-4
3
12`
	testutil.AssertEqualCase(t, rawText, 0, CF1554B)
}
