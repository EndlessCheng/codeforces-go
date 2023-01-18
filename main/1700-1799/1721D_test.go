package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1721/D
// https://codeforces.com/problemset/status/1721/problem/D
func TestCF1721D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5
1 0 0 3 3
2 3 2 1 0
3
1 1 1
0 0 3
8
0 1 2 3 4 5 6 7
7 6 5 4 3 2 1 0
outputCopy
2
0
7
inputCopy
1
8
28 14 12 27 10 8 27 27
5 23 17 2 21 19 6 22
outputCopy
12
inputCopy
1
3
27 30 0
0 5 9
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, -1, CF1721D)
}
