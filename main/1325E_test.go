package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 4 6
outputCopy
1
inputCopy
4
2 3 6 6
outputCopy
2
inputCopy
3
6 15 10
outputCopy
3
inputCopy
4
2 3 5 7
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1325E)
}
