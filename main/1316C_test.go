package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1316C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 2
1 1 2
2 1
outputCopy
1
inputCopy
2 2 999999937
2 1
3 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1316C)
}
