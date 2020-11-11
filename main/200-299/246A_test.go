package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF246A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF246A)
}
