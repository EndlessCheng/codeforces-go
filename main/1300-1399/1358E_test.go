package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 -1
2
outputCopy
2
inputCopy
5
2 2 -8
2
outputCopy
-1
inputCopy
6
-2 -2 6
-1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1358E)
}
