package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183H(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
asdf
outputCopy
4
inputCopy
5 6
aaaaa
outputCopy
15
inputCopy
5 7
aaaaa
outputCopy
-1
inputCopy
10 100
ajihiushda
outputCopy
233`
	testutil.AssertEqualCase(t, rawText, 0, CF1183H)
}
