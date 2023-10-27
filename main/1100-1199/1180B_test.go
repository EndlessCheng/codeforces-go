package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1180/B
// https://codeforces.com/problemset/status/1180/problem/B
func TestCF1180B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2 2 2
outputCopy
-3 -3 -3 -3 
inputCopy
1
0
outputCopy
0 
inputCopy
3
-3 -3 2
outputCopy
-3 -3 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1180B)
}
