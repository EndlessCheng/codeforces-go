package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF87C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
2
inputCopy
6
outputCopy
-1
inputCopy
100
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF87C)
}
