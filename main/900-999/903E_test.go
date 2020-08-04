package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF903E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
abac
caab
acba
outputCopy
acab
inputCopy
3 4
kbbu
kbub
ubkb
outputCopy
kbub
inputCopy
5 4
abcd
dcba
acbd
dbca
zzzz
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF903E)
}
