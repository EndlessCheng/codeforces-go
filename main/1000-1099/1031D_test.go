package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1031D(t *testing.T) {
	// just copy from website
	rawText := `
4 2
abcd
bcde
bcad
bcde
outputCopy
aaabcde
inputCopy
5 3
bwwwz
hrhdh
sepsp
sqfaf
ajbvw
outputCopy
aaaepfafw
inputCopy
7 6
ypnxnnp
pnxonpm
nxanpou
xnnpmud
nhtdudu
npmuduh
pmutsnz
outputCopy
aaaaaaadudsnz`
	testutil.AssertEqualCase(t, rawText, 0, CF1031D)
}
