package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1029F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
outputCopy
12
inputCopy
3 9
outputCopy
14
inputCopy
9 3
outputCopy
14
inputCopy
3 6
outputCopy
12
inputCopy
506 2708
outputCopy
3218`
	testutil.AssertEqualCase(t, rawText, 0, CF1029F)
}
