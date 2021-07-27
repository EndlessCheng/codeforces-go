package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1215/E
// https://codeforces.com/problemset/status/1215/problem/E
func TestCF1215E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 4 2 3 4 2 2
outputCopy
3
inputCopy
5
20 1 14 10 2
outputCopy
0
inputCopy
13
5 5 4 4 3 5 7 6 5 4 4 6 5
outputCopy
21`
	testutil.AssertEqualCase(t, rawText, 0, CF1215E)
}
