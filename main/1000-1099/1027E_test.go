package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1027/problem/E
// https://codeforces.com/problemset/status/1027/problem/E
func TestCF1027E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
0
inputCopy
2 3
outputCopy
6
inputCopy
49 1808
outputCopy
359087121`
	testutil.AssertEqualCase(t, rawText, 0, CF1027E)
}
