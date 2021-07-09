package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1030/E
// https://codeforces.com/problemset/status/1030/problem/E
func TestCF1030E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6 7 14
outputCopy
2
inputCopy
4
1 2 1 16
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1030E)
}
