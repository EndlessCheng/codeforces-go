package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/333/problem/B
// https://codeforces.com/problemset/status/333/problem/B
func TestCF333B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
2 2
outputCopy
0
inputCopy
3 0
outputCopy
1
inputCopy
4 3
3 1
3 2
3 3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF333B)
}
