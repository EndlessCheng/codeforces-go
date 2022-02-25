package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1340/C
// https://codeforces.com/problemset/status/1340/problem/C
func TestCF1340C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
15 5
0 3 7 14 15
11 11
outputCopy
45
inputCopy
13 4
0 3 7 13
9 9
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 1, CF1340C)
}
