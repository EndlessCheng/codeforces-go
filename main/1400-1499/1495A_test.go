package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1495/A
// https://codeforces.com/problemset/status/1495/problem/A
func TestCF1495A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
0 1
1 0
0 -1
-2 0
4
1 0
3 0
-5 0
6 0
0 3
0 1
0 2
0 4
5
3 0
0 4
0 -3
4 0
2 0
1 0
-3 0
0 -10
0 -2
0 -10
outputCopy
3.650281539872885
18.061819283610362
32.052255376143336`
	testutil.AssertEqualCase(t, rawText, 0, CF1495A)
}
