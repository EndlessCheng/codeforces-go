package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1355C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 2 3 4
outputCopy
4
inputCopy
1 2 2 5
outputCopy
3
inputCopy
500000 500000 500000 500000
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1355C)
}
