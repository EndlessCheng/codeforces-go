package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// CF1310D

// https://codeforces.com/problemset/problem/1314/D
// https://codeforces.com/problemset/status/1314/problem/D
func TestCF1314D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 8
0 1 2 2 0
0 0 1 1 2
0 1 0 0 0
2 1 1 0 0
2 0 1 2 0
outputCopy
2
inputCopy
3 2
0 1 1
2 0 1
2 2 0
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1314D)
}
