package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/570/B
// https://codeforces.com/problemset/status/570/problem/B
func TestCF570B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
outputCopy
2
inputCopy
4 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF570B)
}
