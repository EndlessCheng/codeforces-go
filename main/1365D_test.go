package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1365D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 1
.
1 2
G.
2 2
#B
G.
2 3
G.#
B#.
3 3
#B.
#..
GG.
2 2
#B
B.
outputCopy
Yes
Yes
No
No
Yes
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1365D)
}
