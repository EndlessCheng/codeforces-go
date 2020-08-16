package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1225C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
24 0
outputCopy
2
inputCopy
24 1
outputCopy
3
inputCopy
24 -1
outputCopy
4
inputCopy
4 -7
outputCopy
2
inputCopy
1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1225C)
}
