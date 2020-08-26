package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1400E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 4 1 1
outputCopy
2
inputCopy
5
1 0 1 0 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1400E)
}
