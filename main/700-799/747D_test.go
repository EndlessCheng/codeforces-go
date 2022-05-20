package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/747/D
// https://codeforces.com/problemset/status/747/problem/D
func TestCF747D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
-5 20 -3 0
outputCopy
2
inputCopy
4 2
-5 20 -3 0
outputCopy
4
inputCopy
10 6
2 -5 1 3 0 0 -4 -3 1 0
outputCopy
3
inputCopy
4 4
-5 20 -3 0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF747D)
}
