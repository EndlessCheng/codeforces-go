package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1152/B
// https://codeforces.com/problemset/status/1152/problem/B
func TestCF1152B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
39
outputCopy
4
5 3 
inputCopy
1
outputCopy
0
inputCopy
7
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1152B)
}
