package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1092E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2
2 3
outputCopy
2
4 2
inputCopy
2 0
outputCopy
1
1 2
inputCopy
3 2
1 3
2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1092E)
}
