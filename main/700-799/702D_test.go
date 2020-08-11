package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF702D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2 1 4 10
outputCopy
14
inputCopy
5 2 1 4 5
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, 0, CF702D)
}
