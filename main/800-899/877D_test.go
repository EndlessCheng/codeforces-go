package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/877/D
// https://codeforces.com/problemset/status/877/problem/D
func TestCF877D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 4
....
###.
....
1 1 3 1
outputCopy
3
inputCopy
3 4 1
....
###.
....
1 1 3 1
outputCopy
8
inputCopy
2 2 1
.#
#.
1 1 2 2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF877D)
}
