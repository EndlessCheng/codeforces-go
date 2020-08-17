package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1341D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 7
0000000
outputCopy
8
inputCopy
2 5
0010010
0010010
outputCopy
97
inputCopy
3 5
0100001
1001001
1010011
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1341D)
}
