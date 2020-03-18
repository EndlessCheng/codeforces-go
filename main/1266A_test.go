package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1266A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
603
006
205
228
1053
0000000000000000000000000000000000000000000000
outputCopy
red
red
cyan
cyan
cyan
red`
	testutil.AssertEqualCase(t, rawText, 0, CF1266A)
}
