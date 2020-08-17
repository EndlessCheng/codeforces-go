package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1355D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 4
outputCopy
YES
4
2
inputCopy
3 4
outputCopy
NO
inputCopy
3 8
outputCopy
YES
2 1 5
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1355D)
}
