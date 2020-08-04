package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF923A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
14
outputCopy
6
inputCopy
20
outputCopy
15
inputCopy
8192
outputCopy
8191`
	testutil.AssertEqualCase(t, rawText, 0, CF923A)
}
