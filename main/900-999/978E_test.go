package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5
2 1 -3
outputCopy
3
inputCopy
2 4
-1 1
outputCopy
4
inputCopy
4 10
2 4 1 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF978E)
}
