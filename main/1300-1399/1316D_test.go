package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1316D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1 1 1
2 2 2 2
outputCopy
VALID
XL
RX
inputCopy
3
-1 -1 -1 -1 -1 -1
-1 -1 2 2 -1 -1
-1 -1 -1 -1 -1 -1
outputCopy
VALID
RRD
UXD
ULL
inputCopy
2
1 2 1 1
-1 -1 -1 -1
outputCopy
INVALID`
	testutil.AssertEqualCase(t, rawText, 0, CF1316D)
}
