package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
24 21 14 10
outputCopy
3
inputCopy
2
500 500
outputCopy
0
inputCopy
3
14 5 1
outputCopy
-1
inputCopy
5
1 3 6 9 12
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF978D)
}
