package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF353D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
MFM
outputCopy
1
inputCopy
MMFF
outputCopy
3
inputCopy
FFMMM
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF353D)
}
