package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/8/C
// https://codeforces.com/problemset/status/8/problem/C
func TestCF8C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 0
2
1 1
-1 1
outputCopy
8
0 1 2 0 
inputCopy
1 1
3
4 3
3 4
0 0
outputCopy
32
0 1 2 0 3 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF8C)
}
