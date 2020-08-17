package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1307C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aaabb
outputCopy
6
inputCopy
usaco
outputCopy
1
inputCopy
lol
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1307C)
}
