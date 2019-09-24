package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1200E(t *testing.T) {
	// just copy from website
	rawText := `
3
aaaa a aaaaaa
outputCopy
aaaaaa
inputCopy
5
9999999 999 999999999999999 99999 9999
outputCopy
999999999999999
inputCopy
5
I want Iwant want want
outputCopy
Iwant
inputCopy
5
I want to order pizza
outputCopy
Iwantorderpizza
inputCopy
5
sample please ease in out
outputCopy
sampleaseinout`
	testutil.AssertEqual(t, rawText, Sol1200E)
}
