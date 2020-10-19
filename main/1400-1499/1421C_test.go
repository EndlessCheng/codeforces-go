package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1421C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abac
outputCopy
2
R 2
R 5
inputCopy
acccc
outputCopy
2
L 4
L 2
inputCopy
hannah
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1421C)
}
