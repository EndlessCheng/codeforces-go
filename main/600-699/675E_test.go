package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/675/E
// https://codeforces.com/problemset/status/675/problem/E
func TestCF675E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 4 4
outputCopy
6
inputCopy
5
2 3 5 5
outputCopy
17
inputCopy
10
2 10 8 7 8 8 10 9 10
outputCopy
63
inputCopy
7
7 3 4 6 6 7
outputCopy
35
inputCopy
9
2 9 7 6 9 7 8 9
outputCopy
52`
	testutil.AssertEqualCase(t, rawText, -1, CF675E)
}
