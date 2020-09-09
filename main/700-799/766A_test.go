package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF766A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abcd
defgh
outputCopy
5
inputCopy
a
a
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF766A)
}
