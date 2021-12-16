package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1613/E
// https://codeforces.com/problemset/status/1613/problem/E
func TestCF1613E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 3
...
.L.
...
4 5
#....
..##L
...#.
.....
1 1
L
1 9
....L..#.
outputCopy
...
.L.
...
#++++
..##L
...#+
...++
L
++++L++#.`
	testutil.AssertEqualCase(t, rawText, 0, CF1613E)
}
