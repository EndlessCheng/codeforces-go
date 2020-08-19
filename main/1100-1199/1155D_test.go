package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1155D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 -2
-3 8 -2 1 -6
outputCopy
22
inputCopy
12 -3
1 3 3 7 1 3 3 7 1 3 3 7
outputCopy
42
inputCopy
5 10
-1 -2 -3 -4 -5
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1155D)
}
