package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF359B(t *testing.T) {
	// just copy from website
	rawText := `
1 0
outputCopy
1 2
inputCopy
2 1
outputCopy
3 2 1 4
inputCopy
4 0
outputCopy
2 7 4 6 1 3 5 8`
	testutil.AssertEqualCase(t, rawText, 0, CF359B)
}
