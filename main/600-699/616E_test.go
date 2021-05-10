package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/616/E
// https://codeforces.com/problemset/status/616/problem/E
func TestCF616E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
outputCopy
4
inputCopy
4 4
outputCopy
1
inputCopy
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF616E)
}
