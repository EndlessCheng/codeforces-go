package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1141/E
// https://codeforces.com/problemset/status/1141/problem/E
func TestCF1141E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1000 6
-100 -200 -300 125 77 -4
outputCopy
9
inputCopy
1000000000000 5
-1 0 0 0 0
outputCopy
4999999999996
inputCopy
10 4
-3 -6 5 4
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1141E)
}
