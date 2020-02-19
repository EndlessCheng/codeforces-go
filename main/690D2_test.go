package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF690D2(t *testing.T) {
	// just copy from website
	rawText := `
5 1
outputCopy
5
inputCopy
2 2
outputCopy
5
inputCopy
3 2
outputCopy
9
inputCopy
11 5
outputCopy
4367
inputCopy
37 63
outputCopy
230574`
	testutil.AssertEqualCase(t, rawText, 0, CF690D2)
}
