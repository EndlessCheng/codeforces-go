package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF459C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 2
outputCopy
1 1 2 
1 2 1 
inputCopy
3 2 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF459C)
}
