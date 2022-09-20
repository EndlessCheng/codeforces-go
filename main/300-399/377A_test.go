package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/377/A
// https://codeforces.com/problemset/status/377/problem/A
func TestCF377A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 2
#..#
..#.
#...
outputCopy
#.X#
X.#.
#...
inputCopy
5 4 5
#...
#.#.
.#..
...#
.#.#
outputCopy
#XXX
#X#.
X#..
...#
.#.#`
	testutil.AssertEqualCase(t, rawText, 0, CF377A)
}
