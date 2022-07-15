package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1081/B
// https://codeforces.com/problemset/status/1081/problem/B
func TestCF1081B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 0 0
outputCopy
Possible
1 1 1 
inputCopy
5
3 3 2 2 2
outputCopy
Possible
1 1 2 2 2 
inputCopy
4
0 1 2 3
outputCopy
Impossible`
	testutil.AssertEqualCase(t, rawText, 0, CF1081B)
}
