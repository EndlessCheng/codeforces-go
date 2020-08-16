package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1263D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
a
b
ab
d
outputCopy
2
inputCopy
3
ab
bc
abc
outputCopy
1
inputCopy
1
codeforces
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1263D)
}
