package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1690/E
// https://codeforces.com/problemset/status/1690/problem/E
func TestCF1690E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6 3
3 2 7 1 4 8
4 3
2 1 5 6
4 12
0 0 0 0
2 1
1 1
6 10
2 0 0 5 9 4
6 5
5 3 8 6 3 2
outputCopy
8
4
0
2
1
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1690E)
}
