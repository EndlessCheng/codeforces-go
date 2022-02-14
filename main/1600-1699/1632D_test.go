package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1632/D
// https://codeforces.com/problemset/status/1632/problem/D
func TestCF1632D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
1
outputCopy
1 
inputCopy
3
1 4 2
outputCopy
1 1 2 
inputCopy
7
2 12 4 8 18 3 6
outputCopy
0 1 1 1 2 2 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1632D)
}
