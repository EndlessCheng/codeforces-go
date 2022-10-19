package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1208/E
// https://codeforces.com/problemset/status/1208/problem/E
func TestCF1208E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
3 2 4 8
2 2 5
2 6 3
outputCopy
10 15 16 
inputCopy
2 2
2 7 8
1 -8
outputCopy
7 8 `
	testutil.AssertEqualCase(t, rawText, 0, CF1208E)
}
