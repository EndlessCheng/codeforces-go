package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/912/B
// https://codeforces.com/problemset/status/912/problem/B
func TestCF912B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
outputCopy
7
inputCopy
6 6
outputCopy
7
inputCopy
415853337373441 52
outputCopy
562949953421311
inputCopy
2 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF912B)
}
