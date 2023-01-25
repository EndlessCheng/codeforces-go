package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/961/problem/D
// https://codeforces.com/problemset/status/961/problem/D
func TestCF961D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 0
0 1
1 1
1 -1
2 2
outputCopy
YES
inputCopy
5
0 0
1 0
2 1
1 1
2 3
outputCopy
NO
inputCopy
5
3 3
6 3
0 0
10 0
-10 0
outputCopy
YES
inputCopy
6
-1 -1
-1 -2
-1 -3
1000000000 1
-1000000000 0
999999999 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF961D)
}
