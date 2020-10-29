package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1151C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 3
outputCopy
7
inputCopy
5 14
outputCopy
105
inputCopy
88005553535 99999999999
outputCopy
761141116`
	testutil.AssertEqualCase(t, rawText, 0, CF1151C)
}
