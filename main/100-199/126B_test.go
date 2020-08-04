package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol126B(t *testing.T) {
	// just copy from website
	rawText := `
aabaaabaaaaab
outputCopy
aab
inputCopy
abcabcdabcabcabc
outputCopy
abcabc
inputCopy
abcabcabc
outputCopy
abc
inputCopy
fixprefixsuffix
outputCopy
fix
inputCopy
abcdabc
outputCopy
Just a legend`
	testutil.AssertEqual(t, rawText, Sol126B)
}
