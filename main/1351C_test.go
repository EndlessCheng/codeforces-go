package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1351C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
NNN
NS
WWEN
WWEE
NWNWS
outputCopy
15
6
16
12
25`
	testutil.AssertEqualCase(t, rawText, 0, CF1351C)
}
