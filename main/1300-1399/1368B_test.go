package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1368B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
codeforces
inputCopy
3
outputCopy
codeforcesss`
	testutil.AssertEqualCase(t, rawText, 0, CF1368B)
}
