package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1437E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
1 2 1 1 3 5 1
3 5
outputCopy
4
inputCopy
3 3
1 3 2
1 2 3
outputCopy
-1
inputCopy
5 0
4 3 1 2 3
outputCopy
2
inputCopy
10 3
1 3 5 6 12 9 8 10 13 15
2 4 9
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1437E)
}
