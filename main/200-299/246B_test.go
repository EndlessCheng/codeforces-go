package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF246B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 1
outputCopy
1
inputCopy
3
1 4 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF246B)
}
