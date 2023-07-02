package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1215/B
// https://codeforces.com/problemset/status/1215/problem/B
func TestCF1215B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 -3 3 -1 1
outputCopy
8 7
inputCopy
10
4 2 -4 3 1 2 -4 3 2 3
outputCopy
28 27
inputCopy
5
-1 -2 -3 -4 -5
outputCopy
9 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1215B)
}
