package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1030C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
73452
outputCopy
YES
inputCopy
4
1248
outputCopy
NO
inputCopy
5
00000
outputCopy
YES
inputCopy
5
00010
outputCopy
NO
inputCopy
5
01010
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1030C)
}
