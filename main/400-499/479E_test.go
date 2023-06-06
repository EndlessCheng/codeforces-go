package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/479/E
// https://codeforces.com/problemset/status/479/problem/E
func TestCF479E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2 4 1
outputCopy
2
inputCopy
5 2 4 2
outputCopy
2
inputCopy
5 3 4 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF479E)
}
