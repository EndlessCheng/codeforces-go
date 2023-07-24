package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/746/D
// https://codeforces.com/problemset/status/746/problem/D
func TestCF746D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 3 2
outputCopy
GBGBG
inputCopy
7 2 2 5
outputCopy
BBGBGBB
inputCopy
4 3 4 0
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF746D)
}
