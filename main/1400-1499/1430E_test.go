package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1430E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
aaaza
outputCopy
2
inputCopy
6
cbaabc
outputCopy
0
inputCopy
9
icpcsguru
outputCopy
30`
	testutil.AssertEqualCase(t, rawText, 0, CF1430E)
}
