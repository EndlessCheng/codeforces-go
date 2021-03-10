package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1495/problem/B
// https://codeforces.com/problemset/status/1495/problem/B
func TestCF1495B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 5 4 3
outputCopy
1
inputCopy
7
1 2 4 6 5 3 7
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1495B)
}
