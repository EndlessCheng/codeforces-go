package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/875/D
// https://codeforces.com/problemset/status/875/problem/D
func TestCF875D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2 1 6 5
outputCopy
8
inputCopy
4
3 3 3 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF875D)
}
