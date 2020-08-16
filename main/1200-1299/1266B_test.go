package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1266B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
29 34 19 38
outputCopy
YES
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1266B)
}
