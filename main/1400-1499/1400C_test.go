package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1400C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
101110
2
01
1
110
1
outputCopy
111011
10
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1400C)
}
