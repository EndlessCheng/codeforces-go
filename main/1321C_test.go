package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1321C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
bacabcab
outputCopy
4
inputCopy
4
bcda
outputCopy
3
inputCopy
6
abbbbb
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1321C)
}
