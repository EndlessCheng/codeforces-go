package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF280C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
1.50000000000000000000
inputCopy
3
1 2
1 3
outputCopy
2.00000000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF280C)
}
