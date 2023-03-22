package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1626/D
// https://codeforces.com/problemset/status/1626/problem/D
func TestCF1626D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
3 1 2 1
1
1
6
2 2 2 1 1 1
8
6 3 6 3 6 3 6 6
outputCopy
0
2
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1626D)
}
