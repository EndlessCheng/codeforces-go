package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 1
outputCopy
1100
inputCopy
3 3 3
outputCopy
101100
inputCopy
5 3 6
outputCopy
01010100`
	testutil.AssertEqualCase(t, rawText, 0, CF1003B)
}
