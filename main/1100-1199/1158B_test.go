package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1158/B
// https://codeforces.com/problemset/status/1158/problem/B
func TestCF1158B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
outputCopy
1111
inputCopy
5 3
outputCopy
01010
inputCopy
7 3
outputCopy
1011011
inputCopy
9 1
outputCopy

inputCopy
10 2
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1158B)
}
