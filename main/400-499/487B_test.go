package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/487/B
// https://codeforces.com/problemset/status/487/problem/B
func TestCF487B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2 2
1 3 1 2 4 1 2
outputCopy
3
inputCopy
7 2 2
1 100 1 100 1 100 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF487B)
}
