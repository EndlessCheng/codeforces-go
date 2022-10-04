package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/16/E
// https://codeforces.com/problemset/status/16/problem/E
func TestCF16E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
0 0.5
0.5 0
outputCopy
0.500000 0.500000 
inputCopy
5
0 1 1 1 1
0 0 0.5 0.5 0.5
0 0.5 0 0.5 0.5
0 0.5 0.5 0 0.5
0 0.5 0.5 0.5 0
outputCopy
1.000000 0.000000 0.000000 0.000000 0.000000 `
	testutil.AssertEqualCase(t, rawText, 0, CF16E)
}
