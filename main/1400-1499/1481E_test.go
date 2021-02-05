package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1481/problem/E
// https://codeforces.com/problemset/status/1481/problem/E
func TestCF1481E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 2 1 3
outputCopy
2
inputCopy
5
1 2 2 1 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1481E)
}
