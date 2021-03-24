package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1185/C2
// https://codeforces.com/problemset/status/1185/problem/C2
func TestCF1185C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 15
1 2 3 4 5 6 7
outputCopy
0 0 0 0 0 2 3 
inputCopy
5 100
80 40 40 40 60
outputCopy
0 1 1 2 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1185C2)
}
