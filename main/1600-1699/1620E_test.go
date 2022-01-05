package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1620/E
// https://codeforces.com/problemset/status/1620/problem/E
func TestCF1620E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 3
1 1
2 1 2
1 2
1 1
1 2
2 1 3
outputCopy
3 2 2 3 2 
inputCopy
4
1 1
1 2
1 1
2 2 2
outputCopy
1 2 1 
inputCopy
8
2 1 4
1 1
1 4
1 2
2 2 4
2 4 3
1 2
2 2 7
outputCopy
1 3 3 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1620E)
}
