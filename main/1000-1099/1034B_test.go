package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1034B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
0
inputCopy
3 3
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1034B)
}
