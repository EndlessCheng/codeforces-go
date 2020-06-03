package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF924C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
0 1 0 3 0 2
outputCopy
6
inputCopy
5
0 1 2 1 2
outputCopy
1
inputCopy
5
0 1 1 2 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF924C)
}
