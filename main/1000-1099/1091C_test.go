package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1091C(t *testing.T) {
	// just copy from website
	rawText := `
6
outputCopy
1 5 9 21
inputCopy
16
outputCopy
1 10 28 64 136`
	testutil.AssertEqualCase(t, rawText, 0, CF1091C)
}
