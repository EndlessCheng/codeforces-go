package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF877B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abba
outputCopy
4
inputCopy
bab
outputCopy
2
inputCopy
a
outputCopy
1
inputCopy
b
outputCopy
1
inputCopy
bbabbbaabbbb
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF877B)
}
