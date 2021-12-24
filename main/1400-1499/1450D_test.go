package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1450/D
// https://codeforces.com/problemset/status/1450/problem/D
func TestCF1450D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 5 3 4 2
4
1 3 2 1
5
1 3 3 3 2
10
1 2 3 4 5 6 7 8 9 10
3
3 3 2
outputCopy
10111
0001
00111
1111111111
000`
	testutil.AssertEqualCase(t, rawText, 0, CF1450D)
}
