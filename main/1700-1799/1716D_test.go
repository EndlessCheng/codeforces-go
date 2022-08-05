package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1716/D
// https://codeforces.com/problemset/status/1716/problem/D
func TestCF1716D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 1
outputCopy
1 1 2 2 3 4 5 6 
inputCopy
10 2
outputCopy
0 1 0 1 1 1 1 2 2 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1716D)
}
