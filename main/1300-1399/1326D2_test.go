package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1326D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
a
abcdfdcecba
abbaxyzyx
codeforces
acbba
outputCopy
a
abcdfdcba
xyzyx
c
abba`
	testutil.AssertEqualCase(t, rawText, 0, CF1326D2)
}
