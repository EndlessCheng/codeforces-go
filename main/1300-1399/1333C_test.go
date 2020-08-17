package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1333C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 -3
outputCopy
5
inputCopy
3
41 -41 41
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1333C)
}
