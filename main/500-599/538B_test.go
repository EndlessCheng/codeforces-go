package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/538/B
// https://codeforces.com/problemset/status/538/problem/B
func TestCF538B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
outputCopy
9
1 1 1 1 1 1 1 1 1 
inputCopy
32
outputCopy
3
10 11 11 `
	testutil.AssertEqualCase(t, rawText, 0, CF538B)
}
