package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1254B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 8 5
outputCopy
9
inputCopy
5
3 10 2 1 5
outputCopy
2
inputCopy
4
0 5 15 10
outputCopy
0
inputCopy
1
1
outputCopy
-1
inputCopy
15
1 1 1 0 0 0 1 1 1 0 0 0 1 1 1
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, -1, CF1254B2)
}
