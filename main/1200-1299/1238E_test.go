package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1238/E
// https://codeforces.com/problemset/status/1238/problem/E
func TestCF1238E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
aacabc
outputCopy
5
inputCopy
6 4
aaaaaa
outputCopy
0
inputCopy
15 4
abacabadabacaba
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1238E)
}
