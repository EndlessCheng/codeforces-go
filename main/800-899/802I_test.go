package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF802I(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
aa
abcd
ccc
abcc
outputCopy
5
10
14
12`
	testutil.AssertEqualCase(t, rawText, 0, CF802I)
}