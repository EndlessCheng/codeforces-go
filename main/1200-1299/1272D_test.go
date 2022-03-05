package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1272/D
// https://codeforces.com/problemset/status/1272/problem/D
func TestCF1272D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 5 3 4
outputCopy
4
inputCopy
2
1 2
outputCopy
2
inputCopy
7
6 5 4 3 2 4 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1272D)
}
