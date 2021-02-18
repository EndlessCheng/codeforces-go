package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1487/problem/C
// https://codeforces.com/problemset/status/1487/problem/C
func TestCF1487C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2
3
outputCopy
0 
1 -1 1 
inputCopy
1
5
outputCopy
`
	testutil.AssertEqualCase(t, rawText, -1, CF1487C)
}
