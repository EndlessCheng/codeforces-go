package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1117D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
outputCopy
5
inputCopy
3 2
outputCopy
3
inputCopy
1000000000000000000 3
outputCopy
615472476`
	testutil.AssertEqualCase(t, rawText, 0, CF1117D)
}
