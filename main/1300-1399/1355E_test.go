package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1355E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1 100 100
1 3 8
outputCopy
12
inputCopy
3 100 1 100
1 3 8
outputCopy
9
inputCopy
3 100 100 1
1 3 8
outputCopy
4
inputCopy
5 1 2 4
5 5 3 6 5
outputCopy
4
inputCopy
5 1 2 2
5 5 3 6 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1355E)
}
