package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1132F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
abaca
outputCopy
3
inputCopy
8
abcddcba
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1132F)
}
