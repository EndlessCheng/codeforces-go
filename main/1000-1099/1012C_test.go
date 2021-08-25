package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1012/C
// https://codeforces.com/problemset/status/1012/problem/C
func TestCF1012C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1 1 1 1
outputCopy
1 2 2 
inputCopy
3
1 2 3
outputCopy
0 2 
inputCopy
5
1 2 3 2 2
outputCopy
0 1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1012C)
}
