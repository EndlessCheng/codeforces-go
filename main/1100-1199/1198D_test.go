package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1198/D
// https://codeforces.com/problemset/status/1198/problem/D
func TestCF1198D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
###
#.#
###
outputCopy
3
inputCopy
3
...
...
...
outputCopy
0
inputCopy
4
#...
....
....
#...
outputCopy
2
inputCopy
5
#...#
.#.#.
.....
.#...
#....
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1198D)
}
