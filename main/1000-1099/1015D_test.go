package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1015/problem/D
// https://codeforces.com/problemset/status/1015/problem/D
func TestCF1015D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 2 15
outputCopy
YES
10 4 
inputCopy
10 9 45
outputCopy
YES
10 1 10 1 2 1 2 1 6 
inputCopy
10 9 81
outputCopy
YES
10 1 10 1 10 1 10 1 10 
inputCopy
10 9 82
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1015D)
}
