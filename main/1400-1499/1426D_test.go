package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 -5 3 2
outputCopy
1
inputCopy
5
4 -2 3 -9 2
outputCopy
0
inputCopy
9
-1 1 -1 1 -1 1 1 -1 -1
outputCopy
6
inputCopy
8
16 -5 -11 -15 10 5 4 -4
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1426D)
}
