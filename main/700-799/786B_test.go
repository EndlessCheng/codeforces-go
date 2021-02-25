package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/786/B
// https://codeforces.com/problemset/status/786/problem/B
func TestCF786B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5 1
2 3 2 3 17
2 3 2 2 16
2 2 2 3 3
3 3 1 1 12
1 3 3 17
outputCopy
0 28 12 
inputCopy
4 3 1
3 4 1 3 12
2 2 3 4 10
1 2 4 16
outputCopy
0 -1 -1 12 `
	testutil.AssertEqualCase(t, rawText, 0, CF786B)
}
