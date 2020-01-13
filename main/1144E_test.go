package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1144E(t *testing.T) {
	// just copy from website
	rawText := `
2
az
bf
outputCopy
bc
inputCopy
5
afogk
asdji
outputCopy
alvuw
inputCopy
6
nijfvj
tvqhwp
outputCopy
qoztvz`
	testutil.AssertEqualCase(t, rawText, 0, CF1144E)
}
