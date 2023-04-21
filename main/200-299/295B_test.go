package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/295/B
// https://codeforces.com/problemset/status/295/problem/B
func TestCF295B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
0
1
outputCopy
0 
inputCopy
2
0 5
4 0
1 2
outputCopy
9 0 
inputCopy
4
0 3 1 1
6 0 400 1
2 4 0 1
1 1 1 0
4 1 2 3
outputCopy
17 23 404 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF295B)
}
