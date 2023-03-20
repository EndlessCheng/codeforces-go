package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1157/E
// https://codeforces.com/problemset/status/1157/problem/E
func TestCF1157E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 1 2 1
3 2 1 1
outputCopy
1 0 0 2 
inputCopy
7
2 5 1 5 3 4 3
2 4 3 5 6 5 1
outputCopy
0 0 0 1 0 2 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF1157E)
}
