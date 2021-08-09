package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/731/E
// https://codeforces.com/problemset/status/731/problem/E
func TestCF731E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 4 8
outputCopy
14
inputCopy
4
1 -7 -2 3
outputCopy
-3`
	testutil.AssertEqualCase(t, rawText, 0, CF731E)
}
