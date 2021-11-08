package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/478/B
// https://codeforces.com/problemset/status/478/problem/B
func TestCF478B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1
outputCopy
10 10
inputCopy
3 2
outputCopy
1 1
inputCopy
6 3
outputCopy
3 6`
	testutil.AssertEqualCase(t, rawText, 0, CF478B)
}
