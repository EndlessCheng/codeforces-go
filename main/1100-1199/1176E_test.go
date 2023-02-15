package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1176/problem/E
// https://codeforces.com/problemset/status/1176/problem/E
func TestCF1176E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 6
1 2
1 3
1 4
2 3
2 4
3 4
6 8
2 5
5 4
4 3
4 1
1 3
2 3
2 6
5 6
outputCopy
2
1 3
3
4 3 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1176E)
}
