package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF687B(t *testing.T) {
	// just copy from website
	rawText := `
4 5
2 3 5 12
outputCopy
Yes
inputCopy
2 7
2 3
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF687B)
}
