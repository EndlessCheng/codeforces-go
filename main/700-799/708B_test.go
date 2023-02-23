package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/708/B
// https://codeforces.com/problemset/status/708/problem/B
func TestCF708B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 2 3 4
outputCopy
Impossible
inputCopy
1 2 2 1
outputCopy
0110
inputCopy
3 8 4 6
outputCopy
1001011`
	testutil.AssertEqualCase(t, rawText, 0, CF708B)
}
