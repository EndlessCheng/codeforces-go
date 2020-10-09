package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1408F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
1
1 2
inputCopy
4
outputCopy
2
1 2
3 4
inputCopy
7
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1408F)
}
