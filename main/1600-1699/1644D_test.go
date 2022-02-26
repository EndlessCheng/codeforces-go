package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1644/problem/D
// https://codeforces.com/problemset/status/1644/problem/D
func TestCF1644D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1 3 2
1 1
1 1
2 2 2 3
2 1
1 1
2 2
outputCopy
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1644D)
}
