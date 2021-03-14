package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/21/B
// https://codeforces.com/problemset/status/21/problem/B
func TestCF21B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 0
2 2 0
outputCopy
-1
inputCopy
1 1 0
2 -2 0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF21B)
}
