package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1461F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2 0
+-*
outputCopy
2*2-0
inputCopy
4
2 1 1 2
+*
outputCopy
2+1+1+2
inputCopy
3
1 2 3
+
outputCopy
1+2+3
inputCopy
3
2 0 2
*-
outputCopy
2-0*2`
	testutil.AssertEqualCase(t, rawText, 0, CF1461F)
}
