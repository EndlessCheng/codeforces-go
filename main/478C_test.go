package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF478C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4 3
outputCopy
4
inputCopy
1 1 1
outputCopy
1
inputCopy
2 3 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF478C)
}
