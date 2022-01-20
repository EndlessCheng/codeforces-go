package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1616/D
// https://codeforces.com/problemset/status/1616/problem/D
func TestCF1616D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 2 3 4 5
2
10
2 4 2 4 2 4 2 4 2 4
3
3
-10 -5 -10
-8
3
9 9 -3
5
outputCopy
4
8
2
2
inputCopy
1
10
5 -9 -1 6 -6 5 -6 -8 5 3
0
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1616D)
}
