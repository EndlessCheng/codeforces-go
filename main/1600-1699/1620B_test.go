package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1620/B
// https://codeforces.com/problemset/status/1620/problem/B
func TestCF1620B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 8
2 1 2
3 2 3 4
3 1 4 6
2 4 5
10 7
2 3 9
2 1 7
3 1 3 4
3 4 5 6
11 5
3 1 6 8
3 3 6 8
3 1 3 4
2 2 4
outputCopy
25
42
35`
	testutil.AssertEqualCase(t, rawText, 0, CF1620B)
}
