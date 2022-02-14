package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1144/D
// https://codeforces.com/problemset/status/1144/problem/D
func TestCF1144D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 4 6 6 6
outputCopy
2
1 2 3 
1 1 2 
inputCopy
3
2 8 10
outputCopy
2
2 2 1 
2 3 2 
inputCopy
4
1 1 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1144D)
}
