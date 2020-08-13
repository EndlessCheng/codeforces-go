package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF895C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1 1
outputCopy
15
inputCopy
4
2 2 2 2
outputCopy
7
inputCopy
5
1 2 4 5 8
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF895C)
}