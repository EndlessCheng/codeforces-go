package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
3 4 1 2
outputCopy
2.666666666666667`
	testutil.AssertEqualCase(t, rawText, 0, CF1003C)
}
