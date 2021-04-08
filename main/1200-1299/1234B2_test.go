package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1234/problem/B2
// https://codeforces.com/problemset/status/1234/problem/B2
func TestCF1234B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
1 2 3 2 1 3 2
outputCopy
2
2 1 
inputCopy
10 4
2 3 3 1 1 2 1 2 3 3
outputCopy
3
1 3 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1234B2)
}
