package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1076B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
outputCopy
1
inputCopy
4
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1076B)
}
