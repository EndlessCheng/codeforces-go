package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/580/problem/D
// https://codeforces.com/problemset/status/580/problem/D
func TestCF580D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 1
1 1
2 1 1
outputCopy
3
inputCopy
4 3 2
1 2 3 4
2 1 5
3 4 2
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, 0, CF580D)
}
