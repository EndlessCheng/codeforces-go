package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
aaab
aba
outputCopy
8
inputCopy
7 3
abacaba
abc
outputCopy
15
inputCopy
4 3
abca
cba
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
