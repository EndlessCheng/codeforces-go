package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol225C(t *testing.T) {
	// just copy from website
	rawText := `
6 5 1 2
##.#.
.###.
###..
#...#
.##.#
###..
outputCopy
11
inputCopy
2 5 1 1
#####
.....
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, Sol225C)
}
