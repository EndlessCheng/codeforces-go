package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/540/C
// https://codeforces.com/problemset/status/540/problem/C
func TestCF540C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6
X...XX
...XX.
.X..X.
......
1 6
2 2
outputCopy
YES
inputCopy
5 4
.X..
...X
X.X.
....
.XX.
5 3
1 1
outputCopy
NO
inputCopy
4 7
..X.XX.
.XX..X.
X...X..
X......
2 2
1 6
outputCopy
YES
inputCopy
1 1
X
1 1
1 1
outputCopy
NO
inputCopy
3 4
XX.X
X...
.X.X
1 2
1 1
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF540C)
}
