package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1030D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3 3
outputCopy
YES
1 0
2 3
4 1
inputCopy
4 4 7
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1030D)
}
