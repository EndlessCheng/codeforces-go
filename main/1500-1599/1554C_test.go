package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1554/problem/C
// https://codeforces.com/problemset/status/1554/problem/C
func TestCF1554C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 5
4 6
3 2
69 696
123456 654321
outputCopy
4
3
0
640
530866`
	testutil.AssertEqualCase(t, rawText, 0, CF1554C)
}
