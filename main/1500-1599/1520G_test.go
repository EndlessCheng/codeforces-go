package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1520/G
// https://codeforces.com/problemset/status/1520/problem/G
func TestCF1520G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 1
0 -1 0 1 -1
0 20 0 0 -1
-1 -1 -1 -1 -1
3 0 0 0 0
-1 0 0 0 0
outputCopy
14`
	testutil.AssertEqualCase(t, rawText, 0, CF1520G)
}
