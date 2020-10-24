package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1268A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
353
outputCopy
3
353
inputCopy
4 2
1234
outputCopy
4
1313`
	testutil.AssertEqualCase(t, rawText, 0, CF1268A)
}
