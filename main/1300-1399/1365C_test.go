package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1365C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
2 3 4 5 1
outputCopy
5
inputCopy
5
5 4 3 2 1
1 2 3 4 5
outputCopy
1
inputCopy
4
1 3 2 4
4 2 3 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1365C)
}
