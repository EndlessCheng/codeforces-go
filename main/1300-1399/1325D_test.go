package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4
outputCopy
2
3 1
inputCopy
1 3
outputCopy
3
1 1 1
inputCopy
8 5
outputCopy
-1
inputCopy
0 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1325D)
}
