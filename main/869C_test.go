package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF869C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1
outputCopy
8
inputCopy
1 2 2
outputCopy
63
inputCopy
1 3 5
outputCopy
3264
inputCopy
6 2 9
outputCopy
813023575`
	testutil.AssertEqualCase(t, rawText, 0, CF869C)
}
