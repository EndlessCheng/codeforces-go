package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1661/problem/D
// https://codeforces.com/problemset/status/1661/problem/D
func TestCF1661D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
5 4 6
outputCopy
5
inputCopy
6 3
1 2 3 2 2 3
outputCopy
3
inputCopy
6 3
1 2 4 1 2 3
outputCopy
3
inputCopy
7 3
50 17 81 25 42 39 96
outputCopy
92`
	testutil.AssertEqualCase(t, rawText, 0, CF1661D)
}
