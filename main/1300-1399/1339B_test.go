package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1339/B
// https://codeforces.com/problemset/status/1339/problem/B
func TestCF1339B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
6
5 -2 4 8 6 5
4
8 1 4 2
outputCopy
5 5 4 6 8 -2
1 2 4 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1339B)
}
