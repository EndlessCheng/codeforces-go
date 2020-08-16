package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1217C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0110
0101
00001000
0001000
outputCopy
4
3
4
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1217C)
}
