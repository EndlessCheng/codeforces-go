package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1513/problem/F
// https://codeforces.com/problemset/status/1513/problem/F
func TestCF1513F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 4 3 2 1
1 2 3 4 5
outputCopy
4
inputCopy
2
1 3
4 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1513F)
}
