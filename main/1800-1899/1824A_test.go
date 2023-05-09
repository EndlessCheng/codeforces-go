package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1824/problem/A
// https://codeforces.com/problemset/status/1824/problem/A
func TestCF1824A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
3 10
5 5 5
4 6
1 -2 -2 1
5 7
-1 -1 4 -2 -2
6 7
5 -2 -2 -2 -2 -2
6 6
-1 1 4 5 -1 4
6 8
-1 -1 -1 3 -1 -2
6 7
5 -1 -2 -2 -2 -2
3 1
-2 -2 1
2 5
5 -2
1 2
-1
outputCopy
1
3
5
6
5
5
5
1
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1824A)
}
