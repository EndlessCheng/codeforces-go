package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1054/B
// https://codeforces.com/problemset/status/1054/problem/B
func TestCF1054B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 1 2 1
outputCopy
-1
inputCopy
3
1 0 1
outputCopy
1
inputCopy
4
0 1 2 239
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1054B)
}
