package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1140/E
// https://codeforces.com/problemset/status/1140/problem/E
func TestCF1140E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
-1 -1
outputCopy
9
inputCopy
5 2
1 -1 -1 1 2
outputCopy
0
inputCopy
5 3
1 -1 -1 1 2
outputCopy
2
inputCopy
4 200000
-1 -1 12345 -1
outputCopy
735945883`
	testutil.AssertEqualCase(t, rawText, 3, CF1140E)
}
