package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/573/B
// https://codeforces.com/problemset/status/573/problem/B
func TestCF573B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 1 4 6 2 2
outputCopy
3
inputCopy
7
3 3 3 1 3 3 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF573B)
}
