package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/382/C
// https://codeforces.com/problemset/status/382/problem/C
func TestCF382C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 1 7
outputCopy
2
-2 10
inputCopy
1
10
outputCopy
-1
inputCopy
4
1 3 5 9
outputCopy
1
7
inputCopy
4
4 3 4 5
outputCopy
0
inputCopy
2
2 4
outputCopy
3
0 3 6
inputCopy
5
1 3 5 9 13
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF382C)
}
