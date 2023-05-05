package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1763/D
// https://codeforces.com/problemset/status/1763/problem/D
func TestCF1763D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 1 3 2 3
3 2 3 3 2
4 3 4 3 1
5 2 5 2 4
5 3 4 5 4
9 3 7 8 6
20 6 15 8 17
outputCopy
0
1
1
1
3
0
4788
inputCopy
1
3 1 2 1 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1763D)
}
