package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF665C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aab
outputCopy
bab
inputCopy
caaab
outputCopy
cabab
inputCopy
zscoder
outputCopy
zscoder`
	testutil.AssertEqualCase(t, rawText, 0, CF665C)
}
