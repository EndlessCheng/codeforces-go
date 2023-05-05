package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/520/B
// https://codeforces.com/problemset/status/520/problem/B
func TestCF520B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6
outputCopy
2
inputCopy
10 1
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF520B)
}
