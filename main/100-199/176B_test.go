package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/176/B
// https://codeforces.com/problemset/status/176/problem/B
func TestCF176B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ab
ab
2
outputCopy
1
inputCopy
ababab
ababab
1
outputCopy
2
inputCopy
ab
ba
2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF176B)
}
