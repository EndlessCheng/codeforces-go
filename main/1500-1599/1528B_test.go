package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1528/B
// https://codeforces.com/problemset/status/1528/problem/B
func TestCF1528B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
1
inputCopy
2
outputCopy
3
inputCopy
3
outputCopy
6
inputCopy
100
outputCopy
688750769`
	testutil.AssertEqualCase(t, rawText, 0, CF1528B)
}
