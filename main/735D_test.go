package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF735D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
2
inputCopy
27
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF735D)
}
