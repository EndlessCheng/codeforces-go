package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1437A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 4
1 2
120 150
outputCopy
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1437A)
}
