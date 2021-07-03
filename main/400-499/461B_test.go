package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/461/B
// https://codeforces.com/problemset/status/461/problem/B
func TestCF461B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 0
0 1 1
outputCopy
2
inputCopy
6
0 1 1 0 4
1 1 0 0 1 0
outputCopy
1
inputCopy
10
0 1 2 1 4 4 4 0 8
0 0 0 1 0 1 1 0 0 1
outputCopy
27`
	testutil.AssertEqualCase(t, rawText, 0, CF461B)
}
