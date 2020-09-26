package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1129B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
outputCopy
4
6 -8 7 -42
inputCopy
612
outputCopy
7
30 -12 -99 123 -2 245 -300`
	testutil.AssertEqualCase(t, rawText, 0, CF1129B)
}
