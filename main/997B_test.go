package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF997B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
4
inputCopy
2
outputCopy
10
inputCopy
10
outputCopy
244`
	testutil.AssertEqualCase(t, rawText, 0, CF997B)
}
