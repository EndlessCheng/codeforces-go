package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1091D(t *testing.T) {
	// just copy from website
	rawText := `
3
outputCopy
9
inputCopy
4
outputCopy
56
inputCopy
10
outputCopy
30052700`
	testutil.AssertEqualCase(t, rawText, 0, CF1091D)
}
