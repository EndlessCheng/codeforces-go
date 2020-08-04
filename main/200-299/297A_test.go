package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF297A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
01011
0110
outputCopy
YES
inputCopy
0011
1110
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF297A)
}
