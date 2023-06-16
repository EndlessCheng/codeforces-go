package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/602/B
// https://codeforces.com/problemset/status/602/problem/B
func TestCF602B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 3 2
outputCopy
4
inputCopy
11
5 4 5 5 6 7 8 8 8 7 6
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF602B)
}
