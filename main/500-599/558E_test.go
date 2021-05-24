package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/558/E
// https://codeforces.com/problemset/status/558/problem/E
func TestCF558E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 5
abacdabcda
7 10 0
5 8 1
1 4 0
3 6 0
7 10 1
outputCopy
cbcaaaabdd
inputCopy
10 1
agjucbvdfk
1 10 1
outputCopy
abcdfgjkuv`
	testutil.AssertEqualCase(t, rawText, 0, CF558E)
}
