package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1172B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
2 4
outputCopy
16
inputCopy
4
1 2
1 3
1 4
outputCopy
24`
	testutil.AssertEqualCase(t, rawText, 2, CF1172B)
}
