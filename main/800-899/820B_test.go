package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/820/B
// https://codeforces.com/problemset/status/820/problem/B
func TestCF820B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 15
outputCopy
1 2 3
inputCopy
4 67
outputCopy
2 1 3
inputCopy
4 68
outputCopy
4 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF820B)
}
