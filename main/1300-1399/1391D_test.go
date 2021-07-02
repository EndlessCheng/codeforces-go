package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1391D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
101
001
110
outputCopy
2
inputCopy
7 15
000100001010010
100111010110001
101101111100100
010000111111010
111010010100001
000011001111101
111111011010011
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1391D)
}
