package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1617/E
// https://codeforces.com/problemset/status/1617/problem/E
func TestCF1617E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 6 7 8 9
outputCopy
2 5 5
inputCopy
2
4 8
outputCopy
1 2 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1617E)
}
