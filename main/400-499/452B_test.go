package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/452/problem/B
// https://codeforces.com/problemset/status/452/problem/B
func TestCF452B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
1 1
0 0
1 0
0 1
inputCopy
0 10
outputCopy
0 1
0 10
0 0
0 9`
	testutil.AssertEqualCase(t, rawText, 0, CF452B)
}
