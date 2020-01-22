package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1152D(t *testing.T) {
	// just copy from website
	rawText := `
1
outputCopy
1
inputCopy
2
outputCopy
3
inputCopy
3
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1152D)
}
