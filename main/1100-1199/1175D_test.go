package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1175D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
-1 -2 5 -4 8
outputCopy
15
inputCopy
7 6
-3 0 -1 -2 -2 -4 -1
outputCopy
-45
inputCopy
4 1
3 -1 6 0
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1175D)
}
