package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1495/C
// https://codeforces.com/problemset/status/1495/problem/C
func TestCF1495C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 3
X.X
...
X.X
4 4
....
.X.X
....
.X.X
5 5
.X...
....X
.X...
.....
X.X.X
1 10
....X.X.X.
2 2
..
..
outputCopy
XXX
..X
XXX
XXXX
.X.X
.X..
.XXX
.X...
.XXXX
.X...
.X...
XXXXX
XXXXXXXXXX
..
..`
	testutil.AssertEqualCase(t, rawText, 0, CF1495C)
}
