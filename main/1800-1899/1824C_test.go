package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1824/problem/C
// https://codeforces.com/problemset/status/1824/problem/C
func TestCF1824C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 5 7 5 8 4
1 2
1 3
1 4
3 5
4 6
outputCopy
3
inputCopy
8
7 10 7 16 19 9 16 11
1 5
4 2
6 5
5 2
7 2
2 3
3 8
outputCopy
3
inputCopy
4
1 2 1 2
1 2
2 3
4 3
outputCopy
0
inputCopy
9
4 3 6 1 5 5 5 2 7
1 2
2 3
4 1
4 5
4 6
4 7
8 1
8 9
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1824C)
}
