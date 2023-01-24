package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/616/problem/D
// https://codeforces.com/problemset/status/616/problem/D
func TestCF616D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2 3 4 5
outputCopy
1 5
inputCopy
9 3
6 5 1 2 3 2 1 4 5
outputCopy
3 7
inputCopy
3 1
1 2 3
outputCopy
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF616D)
}
