package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF691C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
16
outputCopy
1.6E1
inputCopy
01.23400
outputCopy
1.234
inputCopy
.100
outputCopy
1E-1
inputCopy
100.
outputCopy
1E2`
	testutil.AssertEqualCase(t, rawText, 0, CF691C)
}
