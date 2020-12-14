package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF427E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 6
1 2 3
outputCopy
4
inputCopy
5 5
-7 -6 -3 -1 1
outputCopy
16
inputCopy
1 369
0
outputCopy
0
inputCopy
11 2
-375 -108 1336 1453 1598 1892 2804 3732 4291 4588 4822
outputCopy
18716`
	testutil.AssertEqualCase(t, rawText, 0, CF427E)
}
