package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/675/B
// https://codeforces.com/problemset/status/675/problem/B
func TestCF675B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1 1 1 2
outputCopy
2
inputCopy
3 3 1 2 3
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF675B)
}
