package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF862C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
outputCopy
YES
1 2 4 5 7
inputCopy
3 6
outputCopy
YES
1 2 5`
	testutil.AssertEqualCase(t, rawText, 0, CF862C)
}
