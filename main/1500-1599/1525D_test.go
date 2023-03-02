package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1525/D
// https://codeforces.com/problemset/status/1525/problem/D
func TestCF1525D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 0 0 1 0 0 1
outputCopy
3
inputCopy
6
1 1 1 0 0 0
outputCopy
9
inputCopy
5
0 0 0 0 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1525D)
}
