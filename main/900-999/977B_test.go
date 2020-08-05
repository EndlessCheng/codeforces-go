package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
ABACABA
outputCopy
AB
inputCopy
5
ZZZAA
outputCopy
ZZ`
	testutil.AssertEqualCase(t, rawText, 0, CF977B)
}
