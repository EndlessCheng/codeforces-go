package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1333F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
1 
inputCopy
3
outputCopy
1 1 
inputCopy
9
outputCopy
1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1333F)
}
