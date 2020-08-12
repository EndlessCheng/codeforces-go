package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF988E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5071
outputCopy
4
inputCopy
705
outputCopy
1
inputCopy
1241367
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF988E)
}
