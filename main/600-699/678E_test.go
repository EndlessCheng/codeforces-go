package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/678/problem/E
// https://codeforces.com/problemset/status/678/problem/E
func TestCF678E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0.0 0.5 0.8
0.5 0.0 0.4
0.2 0.6 0.0
outputCopy
0.680000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF678E)
}
