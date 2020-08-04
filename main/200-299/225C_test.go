package _00_299

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol225C(t *testing.T) {
	// just copy from website
	rawText := `
5 10 4 16
.#####....
##..#..##.
.#..##.#..
##..#####.
...#.##..#
outputCopy
21
inputCopy
10 5 3 7
.####
###..
##.##
#..#.
.#...
#.##.
.##..
.#.##
#.#..
.#..#
outputCopy
24
inputCopy
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
	testutil.AssertEqualCase(t, rawText, -1, Sol225C)
}
