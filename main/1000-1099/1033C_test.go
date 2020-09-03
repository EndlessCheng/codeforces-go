package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1033C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
3 6 5 4 2 7 1 8
outputCopy
BAAAABAB
inputCopy
15
3 11 2 5 10 9 7 13 15 8 4 12 6 1 14
outputCopy
ABAAAABBBAABAAB`
	testutil.AssertEqualCase(t, rawText, 0, CF1033C)
}
