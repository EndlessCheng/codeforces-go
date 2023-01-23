package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1661/problem/C
// https://codeforces.com/problemset/status/1661/problem/C
func TestCF1661C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 2 4
5
4 4 3 5 5
7
2 5 4 8 3 7 4
outputCopy
4
3
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1661C)
}
