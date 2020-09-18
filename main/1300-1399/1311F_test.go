package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1311F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3 2
-100 2 3
outputCopy
3
inputCopy
5
2 1 4 3 5
2 2 2 3 4
outputCopy
19
inputCopy
2
2 1
-3 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1311F)
}
