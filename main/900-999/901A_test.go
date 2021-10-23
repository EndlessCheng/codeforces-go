package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/901/A
// https://codeforces.com/problemset/status/901/problem/A
func TestCF901A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1 1
outputCopy
perfect
inputCopy
2
1 2 2
outputCopy
ambiguous
0 1 1 3 3
0 1 1 3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF901A)
}
