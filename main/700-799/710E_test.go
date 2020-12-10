package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF710E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 1 1
outputCopy
4
inputCopy
8 1 10
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF710E)
}
