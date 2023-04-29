package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/786/C
// https://codeforces.com/problemset/status/786/problem/C
func TestCF786C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 3 4 3 3
outputCopy
4 2 1 1 1 
inputCopy
8
1 5 7 8 1 7 6 1
outputCopy
8 4 3 2 1 1 1 1 `
	testutil.AssertEqualCase(t, rawText, -1, CF786C)
}
