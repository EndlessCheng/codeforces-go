package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1214D(t *testing.T) {
	// just copy from website
	rawText := `
2 2
.#
#.
outputCopy
0
inputCopy
2 2
..
..
outputCopy
2
inputCopy
4 4
..#.
..##
....
.#..
outputCopy
1
inputCopy
3 4
....
.##.
....
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, Sol1214D)
}
