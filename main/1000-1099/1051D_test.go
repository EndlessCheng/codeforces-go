package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1051D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
outputCopy
12
inputCopy
4 1
outputCopy
2
inputCopy
1 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1051D)
}
