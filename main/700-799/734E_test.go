package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/734/E
// https://codeforces.com/problemset/status/734/problem/E
func TestCF734E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
11
0 0 0 1 1 0 1 0 0 1 1
1 2
1 3
2 4
2 5
5 6
5 7
3 8
3 9
3 10
9 11
outputCopy
2
inputCopy
4
0 0 0 0
1 2
2 3
3 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF734E)
}
